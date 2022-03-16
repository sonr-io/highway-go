package models

import (
	"github.com/koesie10/webauthn/webauthn"
)

type User struct {
	Did         string
	Jwt         Jwt
	Auths       []webauthn.Authenticator
	webAuthName string
	DisplayName string
	Names       []string
}

// WebAuthID should return the ID of the user. This could for example be the binary encoding of an int.
func (user *User) WebAuthID() []byte {
	return []byte(user.Did)
}

// WebAuthName should return the name of the user.
func (user *User) WebAuthName() string {
	return user.webAuthName
}

// WebAuthDisplayName should return the display name of the user.
func (user *User) WebAuthDisplayName() string {
	return user.DisplayName
}
