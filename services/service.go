package service

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"github.com/koesie10/webauthn/webauthn"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	controller "github.com/sonr-io/highway-go/controllers"
	db "github.com/sonr-io/highway-go/database"
	rt "go.buf.build/grpc/go/sonr-io/sonr/registry"
)

// TODO expand with some kind of middleware later
func AddHandlers(r *mux.Router, ctrl *controller.Controller) {
	// hello handler
	r.HandleFunc("/health", HealthHandler(ctrl)).Methods("GET").Schemes("http")

	// JWT handler - DEPRECATED
	// params:
	// token - encoded jwt
	// siganture - signature to attach to DID
	r.HandleFunc("/jwt/generate/{did}", GenerateJWT(ctrl)).Methods("POST").Schemes("http")

	// Start a new account registeration
	r.HandleFunc("/auth/register/begin/{username}", AuthRegisterBegin(ctrl)).Methods("GET").Schemes("http")

	// Finish an account registeration
	r.HandleFunc("/auth/register/finish/{username}", AuthRegisterFinish(ctrl)).Methods("POST").Schemes("http")

	// Begin login to an existig account
	r.HandleFunc("/auth/login/begin/{username}", AuthLoginBegin(ctrl)).Methods("GET").Schemes("http")

	// Finish logging in to an existing account
	r.HandleFunc("/auth/login/finish/{username}", AuthLoginFinish(ctrl)).Methods("POST").Schemes("http")

	// check name
	r.HandleFunc("/check/name/{name}", CheckName(ctrl)).Methods("GET").Schemes("http")

	// record registered name
	r.HandleFunc("/record/name/{did}", RecordName(ctrl)).Methods("POST").Schemes("http")

	// Register a name
	r.HandleFunc("/register/name/{did}", RegisterName(ctrl)).Methods("POST").Schemes("http")
}

// Error Definitions //TODO this has been used twice, move it a layer back and call it instead
var (
	logger                 = golog.Default.Child("grpc/highway")
	ErrEmptyQueue          = errors.New("no items in Transfer Queue")
	ErrInvalidQuery        = errors.New("no SName or PeerID provided")
	ErrMissingParam        = errors.New("paramater is missing")
	ErrProtocolsNotSet     = errors.New("node Protocol has not been initialized")
	ErrMethodUnimplemented = errors.New("method is not implemented")
)

// Hello Handler
func HealthHandler(ctrl *controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// A very simple health check.
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}
}

//signature is based on face or finger ID
// GenerateJWT generates a JWT for the given SName and PeerID.
func GenerateJWT(ctrl *controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		keys, ok := req.URL.Query()["token"]
		if !ok || len(keys[0]) < 1 {
			logger.Warn("Url Param 'token' is missing")
			return
		}
		token := keys[0]

		keys, ok = req.URL.Query()["signature"]
		if !ok || len(keys[0]) < 1 {
			logger.Warn("Url Param 'signature' is missing")
			return
		}
		signature := keys[0]

		// call ctrl
		result, err := ctrl.GenerateDid(ctx, signature, token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		//w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func AuthRegisterBegin(ctrl *controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		vars := mux.Vars(req)
		name, ok := vars["username"]
		if !ok {
			jsonResponse(w, fmt.Errorf("must supply a valid signature from face or touch ID"), http.StatusBadRequest)
			return
		}

		user := ctrl.FindUserByName(ctx, name)

		// user doesn't exist, create new user
		if user.DisplayName == "" {

			taken, _ := ctrl.CheckName(ctx, name)
			if taken {
				jsonResponse(w, fmt.Errorf("username is not availabel to use"), http.StatusBadRequest)
				return
			}
			var names []string
			names = append(names, name)
			did := "did:sonr:temp" + name
			user.DisplayName = name
			user.Names = names
			user.Did = did

			ctrl.NewUser(ctx, *user)
		}

		// registerOptions := func(credCreationOpts *protocol.PublicKeyCredentialCreationOptions) {
		// 	credCreationOpts.CredentialExcludeList = user.CredentialExcludeList()
		// }

		// Get a session. We're ignoring the error resulted from decoding an
		// existing session: Get() always returns a session, even if empty.
		sess, _ := ctrl.Store.Get(req, name)

		// generate PublicKeyCredentialCreationOptions, session data
		ctrl.WebAuth.StartRegistration(req, w, user, webauthn.WrapMap(sess.Values))

		// store session data as marshaled JSON
		err := sess.Save(req, w)
		if err != nil {
			jsonResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func AuthRegisterFinish(ctrl *controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		vars := mux.Vars(req)
		name := vars["username"]

		//get signature
		//bodyBytes, _ := ioutil.ReadAll(req.Body)

		// var attestationResponse protocol.AttestationResponse
		// d := json.NewDecoder(req.Body)
		// d.DisallowUnknownFields()
		// if err := d.Decode(&attestationResponse); err != nil {
		// 	jsonResponse(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }

		// p, err := protocol.ParseAttestationResponse(attestationResponse)
		// if err != nil {
		// 	jsonResponse(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }

		// // attach signature to did
		// placeHolderDid := "did:sonr:temp" + name
		// signature := p.Response.Attestation.AuthData.AttestedCredentialData.CredentialID
		// ctrl.AttachDid(ctx, placeHolderDid, string(signature))

		// req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// get user
		user := ctrl.FindUserByName(ctx, name)
		// user doesn't exist
		if !contains(user.Names, name) {
			jsonResponse(w, fmt.Errorf("must supply a valid username for account"), http.StatusBadRequest)
			return
		}

		// load the session data
		sess, err := ctrl.Store.Get(req, name)
		if err != nil {
			jsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctrl.WebAuth.FinishRegistration(req, w, user, sess)
	}
}

func AuthLoginBegin(ctrl *controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		vars := mux.Vars(req)
		name := vars["username"]

		// get user
		//did := "did:sonr:" + signature
		user := ctrl.FindUserByName(ctx, name)
		// user doesn't exist
		if !contains(user.Names, name) {
			jsonResponse(w, fmt.Errorf("must supply a valid name for account"), http.StatusBadRequest)
			return
		}

		// Get a session.
		sess, _ := ctrl.Store.Get(req, name)

		// generate PublicKeyCredentialCreationOptions, session data
		ctrl.WebAuth.StartLogin(req, w, user, webauthn.WrapMap(sess.Values))

		// store session data as marshaled JSON
		err := sess.Save(req, w)
		if err != nil {
			jsonResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func AuthLoginFinish(ctrl *controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		vars := mux.Vars(req)
		name := vars["username"]

		// get user
		//did := "did:sonr:" + signature
		user := ctrl.FindDid(ctx, name)
		// user doesn't exist
		if !contains(user.Names, name) {
			jsonResponse(w, fmt.Errorf("must supply a valid username for account"), http.StatusBadRequest)
			return
		}

		// load the session data
		sess, err := ctrl.Store.Get(req, name)
		if err != nil {
			jsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctrl.WebAuth.FinishLogin(req, w, user, sess)
	}
}

type Response struct {
	Available bool
}

func CheckName(ctrl *controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		vars := mux.Vars(req)
		name := vars["name"]
		var err error

		//TODO add error checking for bad values

		start := time.Now()
		e := log.Info()
		defer func(e *zerolog.Event, start time.Time) {
			if err != nil {
				e = log.Error().Stack().Err(err)
			}
			e.Str("handler", "CheckName").AnErr("context", ctx.Err()).Str("name", name).Int64("resp_time", time.Now().Sub(start).Milliseconds()).Send()
		}(e, start)

		nameAvailable, err := ctrl.CheckName(ctx, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//format response
		responseObj := Response{Available: nameAvailable}
		js, err := json.Marshal(responseObj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	}
}

func RecordName(ctrl *controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		var err error

		vars := mux.Vars(req)
		did := vars["did"]

		start := time.Now()
		e := log.Info()
		defer func(e *zerolog.Event, start time.Time) {
			if err != nil {
				e = log.Error().Stack().Err(err)
			}
			e.Str("handler", "RecordName").AnErr("context", ctx.Err()).Int64("resp_time", time.Now().Sub(start).Milliseconds()).Send()
		}(e, start)

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		log.Debug().Str("handler", "recordName").Bytes("request_body", body).Send()

		var recObj db.RecordNameObj
		err = json.Unmarshal(body, &recObj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		available, err := ctrl.CheckName(ctx, recObj.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if !available {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = ctrl.InsertRecord(ctx, recObj, did)

		if err != nil {
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func RegisterName(ctrl *controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//var body *rt.MsgRegisterName
		ctx := req.Context()
		var err error

		vars := mux.Vars(req)
		did := vars["did"]

		start := time.Now()
		e := log.Info()
		defer func(e *zerolog.Event, start time.Time) {
			if err != nil {
				e = log.Error().Stack().Err(err)
			}
			e.Str("handler", "RegisterName").AnErr("context", ctx.Err()).Int64("resp_time", time.Now().Sub(start).Milliseconds()).Send()
		}(e, start)

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		log.Debug().Str("handler", "RegisterName").Bytes("request_body", body).Send()

		var recObj *rt.MsgRegisterName
		err = json.Unmarshal(body, &recObj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		resp, err := ctrl.RegisterName(ctx, recObj, did)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		//format response
		js, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

// from: https://github.com/duo-labs/webauthn.io/blob/3f03b482d21476f6b9fb82b2bf1458ff61a61d41/server/response.go#L15
// TODO switch all repsonses like this
func jsonResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.Marshal(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}

// randToken generates a random hex value.
func randToken(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
