package controller

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/kataras/jwt"
	"github.com/koesie10/webauthn/webauthn"
	"github.com/sonr-io/highway-go/config"
	db "github.com/sonr-io/highway-go/database"
	"github.com/sonr-io/highway-go/models"
	"github.com/sonr-io/sonr/x/registry/types"
	rt "go.buf.build/grpc/go/sonr-io/sonr/registry"
)

// any other services required by http server will flow through here
type Controller struct {
	client      db.MongoClient
	privateKey  string
	devAccount  string
	WebAuth     *webauthn.WebAuthn
	highwayStub *models.HighwayStub
}

func New(mongoClient db.MongoClient, cnfg *config.SonrConfig, stub *models.HighwayStub, auth *webauthn.WebAuthn) (*Controller, error) {
	return &Controller{
		client:      mongoClient,
		privateKey:  cnfg.SecretKey,
		devAccount:  cnfg.DevAccount,
		WebAuth:     auth,
		highwayStub: stub,
	}, nil
}

func (ctrl *Controller) CheckName(ctx context.Context, name string) (bool, error) {
	result, err := ctrl.client.CheckName(name)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (ctrl *Controller) InsertRecord(ctx context.Context, recordObj db.RecordNameObj, did string) error {
	successful := ctrl.client.StoreRecord(recordObj, did)

	if !successful {
		return errors.New("mongo error in insert record")
	}

	return nil
}

func (ctrl *Controller) NewUser(ctx context.Context, user models.User) error {
	return ctrl.client.NewUser(user)
}

func (ctrl *Controller) FindUserByName(ctx context.Context, name string) *models.User {
	return ctrl.client.FindUserByName(name)
}

func (ctrl *Controller) FindDid(ctx context.Context, did string) *models.User {
	return ctrl.client.FindDid(did)
}

func (ctrl *Controller) AddCreds(ctx context.Context, user webauthn.User, authenticator webauthn.Authenticator) error {
	return ctrl.client.AddAuthenticator(user, authenticator)
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

func (ctrl *Controller) RegisterName(ctx context.Context, req *rt.MsgRegisterName, did string) (*rt.MsgRegisterNameResponse, error) {
	// account `alice` was initialized during `starport chain serve`
	//accountName := req.Creator
	accountName := ctrl.devAccount // this i shardcoded to the dev account for now //TODO

	// get account from the keyring by account name and return a bech32 address
	address, err := ctrl.highwayStub.Cosmos.Address(accountName)
	if err != nil {
		return &rt.MsgRegisterNameResponse{}, err
	}

	// check for did in db
	fmt.Println(did)
	user := ctrl.client.FindDid(did)
	if user.Did == "" {
		return &rt.MsgRegisterNameResponse{}, errors.New("user does not exist in DB")
	}

	// define a message to create a post
	msg := &types.MsgRegisterName{
		Creator: address.String(),
		//DeviceId:       req.DeviceId,
		NameToRegister: req.NameToRegister,
		//Jwt:            req.PublicKey, //TODO implement new jwk system
	}
	// broadcast a transaction from account accountName with the message to create a post
	//store response in txResp
	txResp, err := ctrl.highwayStub.Cosmos.BroadcastTx(req.Creator, msg)
	if err != nil {
		return &rt.MsgRegisterNameResponse{}, err
	}

	//TODO fix this logic, this is awful
	success := false
	if !txResp.Empty() {
		success = true
	}

	if success {
		ctrl.client.StoreRecord(db.RecordNameObj{Name: req.NameToRegister}, did)
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
