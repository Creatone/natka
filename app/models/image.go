package models

type Image struct {
	ID   string `json:"ID" bson:"_id,omitempty"`
	Data string `json:"data" bson:"data"`
	Type string `json:"type" bson:"type,omitempty"`
}
