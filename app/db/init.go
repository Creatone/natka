package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName      = "natka"
	userCollection    = "users"
	uri               = "mongodb://root:example@localhost:27017"
	connectionTimeout = time.Second * 10
)

var (
	client *mongo.Client
)

func init() {
	client, _ = mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, _ := context.WithTimeout(context.Background(), connectionTimeout)
	err := client.Connect(ctx)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
