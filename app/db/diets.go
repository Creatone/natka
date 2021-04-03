package db

import (
	"go.mongodb.org/mongo-driver/bson"

	"natka/app/models"
)

const dietsCollection = "diets"

func InsertDiet(diet models.Diet) error {
	return insert(dietsCollection, diet)
}

func GetDiets() ([]models.Diet, error) {
	var diets []models.Diet

	err := getAll(dietsCollection, bson.D{}, &diets)
	if err != nil {
		return nil, err
	}

	return diets, nil
}
