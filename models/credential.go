package models

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/duo-labs/webauthn/protocol/webauthncose"

	"github.com/duo-labs/webauthn/webauthn"
	"github.com/jinzhu/gorm"

	log "github.com/sonr-io/highway-go/logger"
)

// Credential is the stored credential for Auth
type Credential struct {
	gorm.Model

	CredentialID string `json:"credential_id"`

	User   User `json:"-"`
	UserID uint `json:"-"`

	Authenticator   Authenticator `json:"authenticator"`
	AuthenticatorID uint          `json:"authenticator_id"`

	PublicKey []byte `json:"public_key,omitempty"`
}

// WebauthnAuthenticator returns the underlying authenticator used to generate
// the credential.
func (c *Credential) WebauthnAuthenticator() webauthn.Authenticator {
	return c.Authenticator.Authenticator
}

func (c *Credential) DisplayPublicKey() string {
	parsedKey, err := webauthncose.ParsePublicKey(c.PublicKey)
	if err != nil {
		log.Error("Error parsing the public key bytes:", err)
		return "Cannot display key"
	}
	switch parsedKey.(type) {
	case webauthncose.RSAPublicKeyData:
		pKey := parsedKey.(webauthncose.RSAPublicKeyData)
		rKey := &rsa.PublicKey{
			N: big.NewInt(0).SetBytes(pKey.Modulus),
			E: int(uint(pKey.Exponent[2]) | uint(pKey.Exponent[1])<<8 | uint(pKey.Exponent[0])<<16),
		}
		data, err := x509.MarshalPKIXPublicKey(rKey)
		if err != nil {
			log.Error("Error marshalling public key to DER:", err)
			return "Cannot display key"
		}
		pemBytes := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: data,
		})
		return fmt.Sprintf("%s", pemBytes)
	case webauthncose.EC2PublicKeyData:
		pKey := parsedKey.(webauthncose.EC2PublicKeyData)
		var curve elliptic.Curve
		switch pKey.Algorithm {
		case -7:
			curve = elliptic.P256()
		case -35:
			curve = elliptic.P384()
		case -36:
			curve = elliptic.P521()
		default:
			log.Error("Error handling curve for EC key")
			return "Cannot display key"
		}
		eKey := &ecdsa.PublicKey{
			Curve: curve,
			X:     big.NewInt(0).SetBytes(pKey.XCoord),
			Y:     big.NewInt(0).SetBytes(pKey.YCoord),
		}
		fmt.Printf("Got formatted key %+v\n", eKey)
		data, err := x509.MarshalPKIXPublicKey(eKey)
		if err != nil {
			log.Error("Error marshalling public key to DER:", err)
			return "Cannot display key"
		}
		pemBytes := pem.EncodeToMemory(&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: data,
		})
		return fmt.Sprintf("%s", pemBytes)
	default:
		return "Cannot display key of this type"
	}
}
