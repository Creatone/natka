package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"

	"github.com/revel/revel"
)

var loginValidationRegex = regexp.MustCompile("^\\w*$")

type User struct {
	ID       *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Name     string              `json:"name" bson:"name"`
	Mail     string              `json:"mail" bson:"mail"`
	Login    string              `json:"login" bson:"login"`
	Password []byte              `json:"password" bson:"password"`
}

func (u *User) String() string {
	return u.Name
}

func (u *User) Validate(v *revel.Validation) {
	v.Check(u.Login,
		revel.Required{},
		revel.MaxSize{Max: 20},
		revel.MinSize{Min: 4},
		revel.Match{Regexp: loginValidationRegex},
	)
}

func (u *User) ValidateMail(v *revel.Validation) {
	v.Check(u.Mail,
		revel.Required{},
		revel.ValidEmail())
}
