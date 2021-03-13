package db

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"natka/app/models"
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
