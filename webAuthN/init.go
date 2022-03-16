package webAuth

import (
	"github.com/koesie10/webauthn/webauthn"
	"github.com/sonr-io/highway-go/config"
	db "github.com/sonr-io/highway-go/database"
)

func New(cnfg *config.SonrConfig, DB *db.MongoClient) (*webauthn.WebAuthn, error) {

	w, err := webauthn.New(&webauthn.Config{
		// A human-readable identifier for the relying party (i.e. your app), intended only for display.
		RelyingPartyName: "webauthn-sonr",
		// Storage for the authenticator.
		AuthenticatorStore: DB,
	})
	if err != nil {
		return nil, err
	}
	return w, nil
}
