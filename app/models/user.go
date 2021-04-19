package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Name     string              `json:"name" bson:"name"`
	Mail     string              `json:"mail" bson:"mail"`
	Password []byte              `json:"password" bson:"password"`
	Admin    bool                `json:"admin" bson:"admin"`
	Avatar   *primitive.ObjectID `json:"avatar_id" bson:"avatar_id"`
}
