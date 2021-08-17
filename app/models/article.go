package models

type Article struct {
	ID          string `json:"ID" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Text        string `json:"text" bson:"text"`
	Thumbnail   string `json:"thumbnail" bson:"thumbnail"`
}
