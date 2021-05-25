package models

type Image struct {
	ID   string `json:"ID" bson:"_id,omitempty"`
	Data []byte `json:"data" bson:"data"`
}
