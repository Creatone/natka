package models

type Diet struct {
	ID          string `json:"ID" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}
