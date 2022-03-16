package db

import (
	"context"
	"time"

	"github.com/koesie10/webauthn/webauthn"
	"github.com/sonr-io/highway-go/models"
	"go.mongodb.org/mongo-driver/bson"
)

// AddAuthenticator should add the given authenticator to a user. The authenticator's type should not be depended
// on; it is constructed by this package. All information should be stored in a way such that it is retrievable
// in the future using GetAuthenticator and GetAuthenticators.
func (db *MongoClient) AddAuthenticator(user webauthn.User, authenticator webauthn.Authenticator) error {
	collection := db.registerColl
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	auth := models.NewWebAuth(authenticator)
	q1 := bson.M{"did": user.WebAuthID()}
	q2 := bson.M{"$addToSet": bson.M{"auths": auth}}
	collection.FindOneAndUpdate(ctx, q1, q2)
	return nil
}

// GetAuthenticator gets a single Authenticator by the given id, as returned by Authenticator.WebAuthID.
func (db *MongoClient) GetAuthenticator(id []byte) (webauthn.Authenticator, error) {
	collection := db.registerColl
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	record := models.Authenticator{}
	collection.FindOne(ctx, bson.M{"auths": id}).Decode(&record)
	return &record, nil
}

// GetAuthenticators gets a list of all registered authenticators for this user. It might be the case that the user
// has been constructed by this package and the only non-empty value is the WebAuthID. In this case, the store
// should still return the authenticators as specified by the ID.
func (db *MongoClient) GetAuthenticators(user webauthn.User) ([]webauthn.Authenticator, error) {
	collection := db.registerColl
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	record := models.User{}
	collection.FindOne(ctx, bson.M{"did": user.WebAuthID()}).Decode(&record)
	return record.Auths, nil
}
