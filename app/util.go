package app

import (
	"github.com/revel/revel/session"

	"natka/app/models"
)

func IsConnected(session session.Session) bool {
	user := &models.User{}
	result, err := session.GetInto("user", user, true)
	if err != nil {
		return false
	}
	if result == nil {
		return false
	}

	return true
}
