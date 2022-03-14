package controller

import (
	"context"

	db "github.com/sonr-io/highway-go/database"
)

// any other services required by http server will flow through here
type Controller struct {
	client     db.MongoClient
	privateKey string
}

func New(client db.MongoClient, key string) *Controller {
	return &Controller{
		client:     client,
		privateKey: key,
	}
}

func (ctrl *Controller) CheckName(ctx context.Context, name string) (bool, error) {
	result, err := ctrl.client.CheckName(name)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (ctrl *Controller) InsertRecord(ctx context.Context, recordObj db.RecordNameObj) error {
	err := ctrl.client.StoreRecord(recordObj)
	if err != nil {
		return err
	}
	return nil
}
