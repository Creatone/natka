package db

import (
	"go.mongodb.org/mongo-driver/bson"

	"natka/app/models"
)

const dietCollection = "diets"

func InsertDiet(diet models.Diet) error {
	return insert(dietCollection, diet)
}

func GetDiets() ([]models.Diet, error) {
	var diets []models.Diet

	err := getAll(dietCollection, bson.D{}, &diets)
	if err != nil {
		return nil, err
	}

	return diets, nil
}
