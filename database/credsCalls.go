package db

import (
	"context"
	"errors"
	"time"

	"github.com/sonr-io/highway-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateCredential creates a new credential object
func (db *MongoClient) CreateCredential(c *models.Credential) error {
	collection := db.creds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.InsertOne(ctx, c)
	return nil
}

// UpdateCredential updates the credential with new attributes.
func (db *MongoClient) UpdateCredential(c *models.Credential) error {
	collection := db.creds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.FindOneAndUpdate(ctx, bson.M{"id": c.ID}, c)
	return nil
}

// GetCredentialsForUser retrieves all credentials for a provided user regardless of relying party.
func (db *MongoClient) GetCredentialsForUser(user *models.User) ([]models.Credential, error) {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.FindOne().SetProjection(bson.D{{"credentials", 1}})
	result := collection.FindOne(ctx, bson.M{"id": user.ID}, opts)
	temp := models.User{}
	err := result.Decode(temp)
	if err != nil {
		return nil, err
	}
	creds := temp.Credentials
	return creds, nil
}

// GetCredentialForUser retrieves a specific credential for a user.
func (db *MongoClient) GetCredentialForUser(user *models.User, credentialID string) (models.Credential, error) {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.FindOne(ctx, bson.M{"id": user.ID}).Decode(user)

	for _, v := range user.Credentials {
		if string(v.ID) == credentialID {
			return v, nil
		}
	}

	return models.Credential{}, errors.New("cred not found on user")
}

// DeleteCredentialByID gets a credential by its ID. In practice, this would be a bad function without
// some other checks (like what user is logged in) because someone could hypothetically delete ANY credential.
func (db *MongoClient) DeleteCredentialByID(credentialID string) error {
	collection := db.creds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.FindOneAndDelete(ctx, bson.M{"id": credentialID})
	return nil
}

// GetCredentialForUser retrieves a specific credential for a user.
func (db *MongoClient) GiveUserCred(username string, cred *models.Credential) error {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	newCred := models.Credential{
		CredentialID:    cred.CredentialID,
		User:            cred.User,
		UserID:          cred.UserID,
		Authenticator:   cred.Authenticator,
		AuthenticatorID: cred.AuthenticatorID,
		PublicKey:       cred.PublicKey,
	}

	collection.FindOneAndUpdate(ctx, bson.M{"username": username}, bson.M{"$push": bson.M{"credentials": newCred}})
	return nil
}
