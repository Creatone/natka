package utils

import (
	"github.com/revel/revel/session"

	"natka/app/models"
)

func IsConnected(session session.Session) *models.User {
	user := &models.User{}
	_, err := session.GetInto("user", user, true)
	if err != nil {
		return nil
	}

	return user
}
