package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID          *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Name        string              `json:"name" bson:"name"`
	Description string              `json:"description" bson:"description"`
	Text        string              `json:"text" bson:"text"`
}
