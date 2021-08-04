package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName      = "natka"
	uri               = "mongodb://root:example@localhost:27017"
	connectionTimeout = time.Second
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

func insert(collectionName string, document interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), connectionTimeout)

	db := client.Database(databaseName)
	col := db.Collection(collectionName)
	if col == nil {
		return nil, errors.New("nil collection")
	}

	result, err := col.InsertOne(ctx, &document)
	if err != nil {
		return nil, err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func edit(collectionName string, filter interface{}, document interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), connectionTimeout)

	db := client.Database(databaseName)
	col := db.Collection(collectionName)
	if col == nil {
		return errors.New("nil collection")
	}

	_, err := col.UpdateOne(ctx, filter, &document)
	if err != nil {
		return err
	}

	return nil
}

func get(collectionName string, filter interface{}, document interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), connectionTimeout)

	db := client.Database(databaseName)
	col := db.Collection(collectionName)

	result := col.FindOne(ctx, filter)
	if result.Err() != nil {
		return result.Err()
	}

	err := result.Decode(document)
	if err != nil {
		return err
	}

	return nil
}

func getAll(collectionName string, filter interface{}, document interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), connectionTimeout)

	db := client.Database(databaseName)
	col := db.Collection(collectionName)

	result, err := col.Find(ctx, filter, &options.FindOptions{
		Sort: bson.D{{Key: "_id", Value: -1}},
	})
	if err != nil {
		return err
	}

	err = result.All(ctx, document)
	if err != nil {
		return err
	}

	return nil
}

func delete(collectionName string, filter interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), connectionTimeout)

	db := client.Database(databaseName)
	col := db.Collection(collectionName)

	result, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return fmt.Errorf("not deleted")
	}

	return nil
}
