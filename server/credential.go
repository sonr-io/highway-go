package server

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"

	"github.com/duo-labs/webauthn/protocol"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/gorilla/mux"
	log "github.com/sonr-io/highway-go/logger"
	"github.com/sonr-io/highway-go/models"
	rt "go.buf.build/grpc/go/sonr-io/sonr/registry"
)

// RequestNewCredential begins a Credential Registration Request, returning a
// PublicKeyCredentialCreationOptions object
func (ws *Server) RequestNewCredential(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var takenNames = []string{"api", "tx", "app", "arianagrande", "azamsharp", "barrybonds", "barrysanders", "billgates", "britneyspears", "cdixon", "cristiano", "drake", "elon", "eminem", "flotus", "iamsrk", "imap", "index", "jack", "jbbernstein", "jeffbezos", "jimmyfallon", "joynerlucas", "jtimberlake", "justinbieber", "katyperry", "kimkardashian", "kingjames", "ladygaga", "larrypage", "launchhouse", "logic", "mail", "main", "markzuckerburg", "meekmill", "naval", "neymarjr", "oprah", "patrickbetdavid", "pop", "potus", "prad", "rihanna", "root", "satyanadella", "sc", "selenagomez", "sergeibrin", "shakira", "shl", "smartrick", "srbachchan", "stephencurry", "sundarpichai", "taylorswift", "tombrady", "vitalik", "michael", "prad2", "papa", "ikj", "ian", "shadowysupercoder", "ianperez", "perez", "0x0", "zac", "smartrick", "holwerda", "zholwerda", "NFT", "classof.o7", "goat", "nsfw", "nick", "ntindle", "nicktindle", "cloud", "devops", "engineer", "ntt", "grace", "get", "gtindle", "0xDEADBEEF", "static", "d0x", "null", "exposure", "zach", "joshLong145", "beanPole", "undefined", "Peyton", "gopher", "cosmic", "lauren", "sonr", "prad", "letsgobrandon", "snr", "erin", "jamey", "monica", "Space", "timmy", "creaton", "Warriors", "BestButt", "Mfers", "Beast", "mary", "david", "RX", "NT", "0X", "OK", "NO", "SN", "GB", "GT", "IP", "AH", "PT", "JL", "AF", "0F", "0p", "00", "C0", "80", "snr", "sonr", "xxxtentacion", "yasht", "teksupport", "luffy", "yeah"}

	username := vars["name"]
	//check length restrictions
	if len(username) < 2 {
		jsonResponse(w, errors.New("name too short"), http.StatusInternalServerError)
		return
	}

	//alphanumeric restrictions
	re := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	if !re.MatchString(username) {
		jsonResponse(w, errors.New("name not alphanumeric"), http.StatusInternalServerError)
		return
	}

	//The trimmer
	if len(username) > 4 && username[len(username)-4:] == ".snr" {
		username = username[:len(username)-4]
	}
	available, err := ws.Ctrl.CheckName(ctx, username)

	// if reserved
	for _, x := range takenNames {
		if strings.ToLower(x) == strings.ToLower(username) {
			available = false
			break
		}
	}

	if err != nil {
		jsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	} else if !available {
		jsonResponse(w, errors.New("name already taken"), http.StatusInternalServerError)
		return
	}

	// Most times relying parties will choose these.
	attType := r.FormValue("attType")
	authType := r.FormValue("authType")

	// Advanced settings
	userVer := r.FormValue("userVerification")
	resKey := r.FormValue("residentKeyRequirement")
	testExtension := r.FormValue("txAuthExtension")

	var residentKeyRequirement *bool
	if strings.EqualFold(resKey, "true") {
		residentKeyRequirement = protocol.ResidentKeyRequired()
	} else {
		residentKeyRequirement = protocol.ResidentKeyUnrequired()
	}

	testEx := protocol.AuthenticationExtensions(map[string]interface{}{"txAuthSimple": testExtension})

	//SQL lite check
	// user, err := models.GetUserByUsername(username)
	// if err != nil {
	// 	user = models.User{
	// 		DisplayName: strings.Split(username, "@")[0],
	// 		Username:    username,
	// 	}
	// 	err = models.PutUser(&user)
	// 	if err != nil {
	// 		jsonResponse(w, "Error creating new user", http.StatusInternalServerError)
	// 		return
	// 	}
	// }

	//secondary mongo check
	user := ws.Ctrl.FindUserByName(ctx, username)

	// user doesn't exist, create new user
	if user.DisplayName == "" {
		available, _ := ws.Ctrl.CheckName(ctx, username)
		if !available {
			jsonResponse(w, fmt.Errorf("username is not availabel to use"), http.StatusAlreadyReported)
			return
		}
		var names []string
		names = append(names, username)
		did := "did:sonr:temp" + username
		user.DisplayName = username
		user.Names = names
		user.Did = did
		user.ID = uint(rand.Uint32())
		user.Username = username
		user.DisplayName = username
		ws.Ctrl.NewUser(ctx, *user)
	}

	credentialOptions, sessionData, err := ws.webauthn.BeginRegistration(user,
		webauthn.WithAuthenticatorSelection(
			protocol.AuthenticatorSelection{
				AuthenticatorAttachment: protocol.AuthenticatorAttachment(authType),
				RequireResidentKey:      residentKeyRequirement,
				UserVerification:        protocol.UserVerificationRequirement(userVer),
			}),
		webauthn.WithConveyancePreference(protocol.ConveyancePreference(attType)),
		webauthn.WithExtensions(testEx),
	)
	if err != nil {
		jsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save the session data as marshaled JSON
	err = ws.store.SaveWebauthnSession("registration", sessionData, r, w)
	if err != nil {
		jsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the PublicKeyCreationOptions back to the browser
	jsonResponse(w, credentialOptions, http.StatusOK)
	return
}

// MakeNewCredential attempts to make a new credential given an authenticator's response
func (ws *Server) MakeNewCredential(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//ctx := r.Context()
	// Load the session data
	sessionData, err := ws.store.GetWebauthnSession("registration", r)
	if err != nil {
		jsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the user associated with the credential
	user, err := ws.Ctrl.GetUser(models.BytesToID(sessionData.UserID))

	if err != nil {
		jsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Verify that the challenge succeeded
	cred, err := ws.webauthn.FinishRegistration(user, sessionData, r)
	if err != nil {
		jsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If needed, you can perform additional checks here to ensure the
	// authenticator and generated credential conform to your requirements.

	// Finally, save the credential and authenticator to the
	// database
	authenticator, err := ws.Ctrl.CreateAuthenticator(cred.Authenticator)
	if err != nil {
		log.Errorf("error creating authenticator: %v", err)
		jsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// For our use case, we're encoding the raw credential ID as URL-safe
	// base64 since we anticipate rendering it in templates. If you choose to
	// do this, make sure to decode the credential ID before passing it back to
	// the webauthn library.
	credentialID := base64.URLEncoding.EncodeToString(cred.ID)
	c := &models.Credential{
		Authenticator:   authenticator,
		AuthenticatorID: authenticator.ID,
		UserID:          user.ID,
		PublicKey:       cred.PublicKey,
		CredentialID:    credentialID,
	}

	fmt.Println(c.CredentialID)
	fmt.Println(c.UserID)

	err = ws.Ctrl.CreateCredential(c)
	if err != nil {
		jsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//store public key on did
	did := "did:sonr:" + credentialID
	regName := &rt.MsgRegisterName{Creator: "", NameToRegister: user.Username}
	ws.Ctrl.AttachDid(ctx, "did:sonr:temp"+user.DisplayName, did)

	//register name on chain

	fmt.Println(did)
	fmt.Println(regName)

	//TODO fix async issue on stripe and no need for cron job
	//ws.Ctrl.RegisterName(ctx, regName, did, c)

	//store cred under user in mgo
	ws.Ctrl.GiveUserCred(user.Username, c)

	jsonResponse(w, http.StatusText(http.StatusCreated), http.StatusCreated)
}

// GetCredentials gets a user's credentials from the db
func (ws *Server) GetCredentials(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//The trimmer
	username := vars["name"]
	if username[len(username)-4:] == ".snr" {
		username = username[:len(username)-4]
	}
	u, err := ws.Ctrl.GetUserByUsername(username)
	if err != nil {
		log.Errorf("user not found: %s: %s", username, err)
		jsonResponse(w, "User not found", http.StatusNotFound)
		return
	}
	cs, err := ws.Ctrl.GetCredentialsForUser(u)
	if err != nil {
		log.Error(err)
		jsonResponse(w, "Credentials not found", http.StatusNotFound)
		return
	}
	jsonResponse(w, cs, http.StatusOK)
}

// DeleteCredential deletes a credential from the db
func (ws *Server) DeleteCredential(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	credID := vars["id"]
	err := ws.Ctrl.DeleteCredentialByID(credID)
	log.Infof("deleting credential: %s", credID)
	if err != nil {
		log.Errorf("error deleting credential: %s", err)
		jsonResponse(w, "Credential not Found", http.StatusNotFound)
		return
	}
	jsonResponse(w, "Success", http.StatusOK)
}

func (ws *Server) HealthHandler(w http.ResponseWriter, req *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

type Response struct {
	Available bool
}
