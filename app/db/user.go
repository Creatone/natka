package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"

	"natka/app/models"
)

const userCollection = "users"

func InsertUser(user models.User) error {
	ctx, _ := context.WithTimeout(context.Background(), connectionTimeout)

	db := client.Database(databaseName)
	col := db.Collection(userCollection)
	if col == nil {
		return errors.New("nil collection")
	}

	_, err := col.InsertOne(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func GetUser(mail string) (*models.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), connectionTimeout)

	db := client.Database(databaseName)
	col := db.Collection(userCollection)

	result := col.FindOne(ctx, bson.D{{"mail", mail}})
	if result.Err() != nil {
		return nil, result.Err()
	}

	user := models.User{}
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
