package models

import "github.com/koesie10/webauthn/webauthn"

type Authenticator struct {
	webAuthID           []byte
	webAuthCredentialID []byte
	webAuthPublicKey    []byte
	webAuthAAGUID       []byte
	webAuthSignCount    uint32
}

func NewWebAuth(authenticator webauthn.Authenticator) Authenticator {
	return Authenticator{
		webAuthID:           authenticator.WebAuthID(),
		webAuthCredentialID: authenticator.WebAuthCredentialID(),
		webAuthPublicKey:    authenticator.WebAuthPublicKey(),
		webAuthAAGUID:       authenticator.WebAuthAAGUID(),
		webAuthSignCount:    authenticator.WebAuthSignCount(),
	}
}

func (auth *Authenticator) WebAuthID() []byte {
	return auth.webAuthID
}
func (auth *Authenticator) WebAuthCredentialID() []byte {
	return auth.webAuthCredentialID
}
func (auth *Authenticator) WebAuthPublicKey() []byte {
	return auth.webAuthPublicKey
}
func (auth *Authenticator) WebAuthAAGUID() []byte {
	return auth.webAuthAAGUID
}
func (auth *Authenticator) WebAuthSignCount() uint32 {
	return auth.webAuthSignCount
}
