package controller

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/duo-labs/webauthn/webauthn"
	"github.com/kataras/jwt"
	"github.com/sonr-io/blockchain/x/registry/types"
	"github.com/sonr-io/highway-go/config"
	db "github.com/sonr-io/highway-go/database"
	"github.com/sonr-io/highway-go/models"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	rt "go.buf.build/grpc/go/sonr-io/sonr/registry"
)

// any other services required by http server will flow through here
type Controller struct {
	client      *db.MongoClient
	privateKey  string
	devAccount  string
	stripeKey   string
	highwayStub *models.HighwayStub
}

func New(mongoClient *db.MongoClient, cnfg *config.SonrConfig, stub *models.HighwayStub) (*Controller, error) {
	return &Controller{
		client:      mongoClient,
		privateKey:  cnfg.SecretKey,
		devAccount:  cnfg.DevAccount,
		highwayStub: stub,
		stripeKey:   cnfg.StripeKey,
	}, nil
}

func (ctrl *Controller) CheckName(ctx context.Context, name string) (bool, error) {
	result, err := ctrl.client.CheckName(name)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (ctrl *Controller) InsertRecord(ctx context.Context, name string, did string) error {
	successful := ctrl.client.StoreRecord(name, did)

	if !successful {
		return errors.New("mongo error in insert record")
	}

	return nil
}

func (ctrl *Controller) NewUser(ctx context.Context, user models.User) error {
	return ctrl.client.NewUser(user)
}

func (ctrl *Controller) GetUser(id uint) (*models.User, error) {
	return ctrl.client.GetUser(id), nil
}

func (ctrl *Controller) FindUserByName(ctx context.Context, name string) *models.User {
	return ctrl.client.FindUserByName(name)
}

func (ctrl *Controller) GetUserByUsername(name string) (*models.User, error) {
	return ctrl.client.GetUserByUsername(name), nil
}

func (ctrl *Controller) PutUser(u *models.User) error {
	ctrl.client.PutUser(u)
	return nil
}

func (ctrl *Controller) GetCredentialsForUser(user *models.User) ([]models.Credential, error) {
	return ctrl.client.GetCredentialsForUser(user)
}

func (ctrl *Controller) CreateAuthenticator(auth webauthn.Authenticator) (models.Authenticator, error) {
	return *ctrl.client.CreateAuthenticator(auth), nil
}

func (ctrl *Controller) CreateCredential(c *models.Credential) error {
	return ctrl.client.CreateCredential(c)
}

func (ctrl *Controller) DeleteCredentialByID(id string) error {
	return ctrl.client.DeleteCredentialByID(id)
}

func (ctrl *Controller) GetCredentialForUser(user models.User, credentialID string) (models.Credential, error) {
	return ctrl.client.GetCredentialForUser(&user, credentialID)
}
func (ctrl *Controller) UpdateAuthenticatorSignCount(id uint, count uint32) error {
	return ctrl.client.UpdateAuthenticatorSignCount(id, count)
}

func (ctrl *Controller) FindDid(ctx context.Context, did string) *models.User {
	return ctrl.client.FindDid(did)
}

func (ctrl *Controller) AttachDid(ctx context.Context, placeHolderDid string, newDid string) error {
	return ctrl.client.AttachDid(placeHolderDid, newDid)
}

func (ctrl *Controller) GiveUserCred(username string, cred *models.Credential) error {
	ctrl.client.GiveUserCred(username, cred)
	return nil
}

// func (ctrl *Controller) AddCreds(ctx context.Context, user webauthn.User, authenticator webauthn.Authenticator) error {
// 	return ctrl.client.AddAuthenticator(user, authenticator)
// }

func (ctrl *Controller) StripeIntent(item models.SnrItem, name string) (*stripe.PaymentIntent, error) {
	stripe.Key = ctrl.stripeKey

	desc := "Payment for the .snr/ name " + name

	// Create a PaymentIntent with amount and currency
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateOrderAmount(item)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
		Description: &desc,
	}

	return paymentintent.New(params)
}

func (ctrl *Controller) AttachIntent(piID string, name string) {
	ctrl.client.AttachIntent(piID, name)
}

func (ctrl *Controller) UpdatePayment(piID string) {
	ctrl.client.SuccessfulPayment(piID)
}

func (ctrl *Controller) RecordPayment(name string) error {
	return ctrl.client.RecordPayment(name)
}

func calculateOrderAmount(item models.SnrItem) int64 {
	// Replace this constant with a calculation of the order's amount
	// Calculate the order total on the server to prevent
	// people from directly manipulating the amount on the client
	return 5000
}

func (ctrl *Controller) GenerateDid(ctx context.Context, signature string, token string) ([]byte, error) {
	verifiedToken, err := jwt.Verify(jwt.HS256, []byte(signature), []byte(token))
	if err != nil {
		return nil, err
	}

	result := models.Jwt{}
	err = verifiedToken.Claims(&result)
	if err != nil {
		return nil, err
	}

	//figure out did
	did := "did:sonr:" + signature

	if ctrl.client.FindDid(did).Did == "" {
		// no record exist make a new one
		ctrl.client.AddDid(did, result)
	}

	return []byte(did), err
}

func (ctrl *Controller) RegisterName(ctx context.Context, req *rt.MsgRegisterName, did string, cred *models.Credential) (*rt.MsgRegisterNameResponse, error) {
	// account `alice` was initialized during `starport chain serve`
	//accountName := req.Creator
	accountName := ctrl.devAccount // this i shardcoded to the dev account for now //TODO

	fmt.Println("account: " + accountName)

	// get account from the keyring by account name and return a bech32 address
	address, err := ctrl.highwayStub.Cosmos.Address(accountName)
	if err != nil {
		fmt.Println(err.Error())
		return &rt.MsgRegisterNameResponse{}, err
	}
	//TODO check for credential

	// check for name in db
	fmt.Println(did)
	user := ctrl.client.FindUserByName(req.NameToRegister)
	if user.DisplayName == "" {
		return &rt.MsgRegisterNameResponse{}, errors.New("user does not exist in DB")
	}
	// define a message to create a post
	msg := &types.MsgRegisterName{
		Creator: address.String(),
		//DeviceId:       req.DeviceId,
		NameToRegister: req.NameToRegister,
		// jwk:idk
	}
	// broadcast a transaction from account accountName with the message to create a post
	//store response in txResp
	txResp, err := ctrl.highwayStub.Cosmos.BroadcastTx(accountName, msg)
	if err != nil {
		fmt.Println(err.Error())
		return &rt.MsgRegisterNameResponse{}, err
	}
	//TODO fix this logic, this is awful
	success := false
	if !txResp.Empty() {
		success = true
	}
	if success {
		ctrl.client.StoreRecord(req.NameToRegister, did)
	}

	// WTF
	// responseTest := types.MsgRegisterNameResponse{}
	// message := txResp.Decode(&responseTest)
	// fmt.Println(message)
	// fmt.Println(&message)
	bs, err := hex.DecodeString(txResp.Data)
	if err != nil {
		return &rt.MsgRegisterNameResponse{}, err
	}

	//Worst case log parsing
	fmt.Println("------------------------------------------------------" + did + "------------------------------------------------------")
	fmt.Println(string(bs))
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	//Unmarshalling of a json did document:
	// parsedDIDDoc := did.Document{}
	// err = json.Unmarshal([]byte(bs), &parsedDIDDoc)
	// if err != nil {
	// 	return &rt.MsgRegisterNameResponse{}, err
	// }
	//TODO unmarshal is not working as intended

	response := rt.MsgRegisterNameResponse{}
	//did := "did:sonr:" + JWT
	response.IsSuccess = success
	response.DidDocumentJson = string(bs)
	response.DidUrl = txResp.TxHash
	// response.DidDocument = &rt.DidDocument{
	// 	AlsoKnownAs: aliases,
	// }
	//return &response, nil
	return &response, nil
}
