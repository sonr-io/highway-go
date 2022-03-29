package db

import (
	"context"
	"time"

	"github.com/duo-labs/webauthn/webauthn"
	"github.com/sonr-io/highway-go/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAuthenticator returns the authenticator the given id corresponds to. If
// no authenticator is found, an error is thrown.
func (db *MongoClient) GetAuthenticator(id uint) *models.Authenticator {
	collection := db.auths
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	auth := &models.Authenticator{}
	collection.FindOne(ctx, bson.M{"id": id}).Decode(auth)
	return auth
}

// CreateAuthenticator creates a new authenticator that's tied to a Credential.
func (db *MongoClient) CreateAuthenticator(a webauthn.Authenticator) *models.Authenticator {
	collection := db.auths
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	auth := &models.Authenticator{}
	auth.Authenticator = a
	collection.InsertOne(ctx, auth)
	return auth
}

// UpdateAuthenticatorSignCount updates a specific authenticator's sign count for tracking
// potential clone attempts.
func (db *MongoClient) UpdateAuthenticatorSignCount(id uint, count uint32) error {
	collection := db.auths
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.FindOneAndUpdate(ctx, bson.M{"id": id}, bson.M{"sign_count": count})
	return nil
}
