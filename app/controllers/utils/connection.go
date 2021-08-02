package utils

import (
	"github.com/revel/revel/session"

	"natka/app/models"
)

const (
	userSessionKey = "user"
)

func IsConnected(session session.Session) *models.User {
	user := &models.User{}
	_, err := session.GetInto("user", user, true)
	if err != nil {
		return nil
	}

	return user
}

func KeepUser(session session.Session, user models.User) error {
	err := session.Set(userSessionKey, user)
	if err != nil {
		return err
	}

	return nil
}
