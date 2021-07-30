package db

import (
	"go.mongodb.org/mongo-driver/bson"

	"natka/app/models"
)

const usersCollection = "users"

func InsertUser(user models.User) (interface{}, error) {
	return insert(usersCollection, user)
}

func GetUser(mail string) (*models.User, error) {
	user := models.User{}

	err := get(usersCollection, bson.D{{"mail", mail}}, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func EditUser(user models.User) error {
	return edit(usersCollection, bson.D{{"_id", user.ID}},
		bson.D{{"$set", bson.D{
			{"name", user.Name},
			{"avatar", user.Avatar},
			{"diets", user.Diets},
		}}})
}
