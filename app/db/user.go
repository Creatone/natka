package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"golang.org/x/crypto/bcrypt"

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

func CheckMail(mail string) (bool, error) {
	_, err := GetUser(mail)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func Login(mail, password string) (*models.User, error) {
	user, err := GetUser(mail)
	if err != nil {
		return nil, err
	}
	if user != nil {
		err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	return nil, nil
}
