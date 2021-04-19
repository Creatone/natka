package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	ID   *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Data []byte              `json:"data" bson:"data"`
}
