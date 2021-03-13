package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Name     string              `json:"name" bson:"name"`
	Mail     string              `json:"mail" bson:"mail"`
	Password []byte              `json:"password" bson:"password"`
}

func (u *User) String() string {
	return u.Name
}
