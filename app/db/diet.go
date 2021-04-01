package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"

	"natka/app/models"
)

const dietCollection = "diets"

func InsertDiet(diet models.Diet) error {
	ctx, _ := context.WithTimeout(context.Background(), connectionTimeout)

	db := client.Database(databaseName)
	col := db.Collection(dietCollection)
	if col == nil {
		return errors.New("nil collection")
	}

	_, err := col.InsertOne(ctx, &diet)
	if err != nil {
		return err
	}
	return nil
}

func GetDiets() ([]models.Diet, error) {
	ctx, _ := context.WithTimeout(context.Background(), connectionTimeout)

	db := client.Database(databaseName)
	col := db.Collection(dietCollection)

	result, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var diets []models.Diet
	err = result.All(ctx, &diets)
	if err != nil {
		return nil, err
	}

	return diets, nil
}
