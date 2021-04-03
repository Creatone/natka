package db

import (
	"go.mongodb.org/mongo-driver/bson"

	"natka/app/models"
)

const userCollection = "users"

func InsertUser(user models.User) error {
	return insert(userCollection, user)
}

func GetUser(mail string) (*models.User, error) {
	user := models.User{}

	err := get(userCollection, bson.D{{"mail", mail}}, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
