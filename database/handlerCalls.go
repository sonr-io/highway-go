package db

import (
	"context"
	"log"
	"time"

	"github.com/sonr-io/highway-go/models"
	"go.mongodb.org/mongo-driver/bson"
)

type RecordNameObj struct {
	Name string `json:"name"`
	// TimeStamp time.Time
}

func (db *MongoClient) FindDid(did string) *models.User {
	collection := db.registerColl
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	record := models.User{}
	collection.FindOne(ctx, bson.M{"did": did}).Decode(&record)
	return &record
}

func (db *MongoClient) AddDid(did string, jwt models.Jwt) error {
	collection := db.registerColl
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	record := &models.User{
		Did: did,
		Jwt: jwt,
	}

	res, err := collection.InsertOne(ctx, record)
	if err != nil || res == nil {
		log.Print("\nunable to insert entry into DB in database package\n")
		return err
	}
	return nil
}

func (db *MongoClient) StoreRecord(recordObj RecordNameObj, did string) bool {
	collection := db.registerColl
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	q1 := bson.M{"did": did}
	q2 := bson.M{"$addToSet": bson.M{"name": recordObj.Name}}

	record := models.User{}
	collection.FindOne(ctx, q1).Decode(&record)

	if record.Did == "" {
		return false
	}

	collection.FindOneAndUpdate(ctx, q1, q2)
	return true
}

// check if name is available, if available return true
func (db *MongoClient) CheckName(name string) (bool, error) {
	collection := db.registerColl
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.CountDocuments(ctx, bson.M{"name": name})
	if err != nil {
		return false, err
	}

	if result == 0 {
		return true, nil
	}
	return false, nil
}

func (db *MongoClient) NewUser(user models.User) error {
	collection := db.registerColl
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.InsertOne(ctx, user)
	return nil
}

func (db *MongoClient) FindUserByName(name string) *models.User {
	collection := db.registerColl
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := models.User{}
	collection.FindOne(ctx, bson.M{"name": name}).Decode(user)
	return &user
}

func (db *MongoClient) AttachDid(placeHolderDid string, newDid string) error {
	collection := db.registerColl
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := models.User{}
	collection.FindOneAndUpdate(ctx, bson.M{"did": placeHolderDid}, bson.M{"$set": bson.M{"did": newDid}}).Decode(user)
	return nil
}
