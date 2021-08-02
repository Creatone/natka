package models

type User struct {
	ID       string              `json:"ID" bson:"_id,omitempty"`
	Name     string              `json:"name" bson:"name"`
	Mail     string              `json:"mail" bson:"mail"`
	Password []byte              `json:"password" bson:"password"`
	Admin    bool                `json:"admin" bson:"admin"`
	Avatar   string              `json:"avatar" bson:"avatar"`
	Diets    map[string]struct{} `json:"diets" bson:"diets"`
}
