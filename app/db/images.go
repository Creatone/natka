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

func EditImage(image models.Image) error {
	return edit(imagesCollection, bson.D{{"_id", image.ID}},
		bson.D{{"$set", bson.D{
			{"data", image.Data}}}})
}
