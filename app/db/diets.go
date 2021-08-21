package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"natka/app/models"
)

const dietsCollection = "diets"

func InsertDiet(diet models.Diet) (interface{}, error) {
	return insert(dietsCollection, diet)
}

func GetDiets() ([]models.Diet, error) {
	var diets []models.Diet

	err := getAll(dietsCollection, bson.D{}, &options.FindOptions{Sort: bson.D{{Key: "_id", Value: -1}}}, &diets)
	if err != nil {
		return nil, err
	}

	return diets, nil
}

func GetDiet(id *primitive.ObjectID) (*models.Diet, error) {
	diet := models.Diet{}

	err := get(dietsCollection, bson.D{{"_id", id}}, &diet)
	if err != nil {
		return nil, err
	}

	return &diet, nil
}

func EditDiet(diet models.Diet) error {
	// TODO: Move bson-key thing to func.
	objectID, err := primitive.ObjectIDFromHex(diet.ID)
	if err != nil {
		return err
	}

	return edit(dietsCollection, bson.D{{"_id", objectID}},
		bson.D{{"$set", bson.D{
			{"name", diet.Name},
			{"description", diet.Description},
		}}})
}

func DeleteDiet(id *primitive.ObjectID) error {
	return remove(dietsCollection, bson.D{{"_id", &id}})
}
