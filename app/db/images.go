package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"natka/app/models"
)

const imagesCollection = "images"

func InsertImage(image models.Image) (interface{}, error) {
	return insert(imagesCollection, image)
}

func GetImage(id string) (models.Image, error) {
	var image models.Image

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return image, err
	}

	err = get(imagesCollection, bson.D{{"_id", objectID}}, &image)
	if err != nil {
		return image, err
	}

	return image, nil
}

func GetImagesByType(passedType string) ([]models.Image, error) {
	var images []models.Image

	err := getAll(imagesCollection, bson.D{{"type", passedType}}, &images)
	if err != nil {
		return images, err
	}

	return images, nil
}

func EditImage(image models.Image) error {
	objectID, err := primitive.ObjectIDFromHex(image.ID)
	if err != nil {
		return err
	}

	return edit(imagesCollection, bson.D{{"_id", objectID}},
		bson.D{{"$set", bson.D{
			{"data", image.Data}}}})
}

func DeleteImage(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return delete(imagesCollection, bson.D{{"_id", &objectID}})
}
