package admin

import (
	"github.com/revel/revel"

	"natka/app/models"
	"natka/app/routes"
)

type Admin struct {
	*revel.Controller
}

func (c Admin) connected() *models.User {
	user := &models.User{}
	_, err := c.Session.GetInto("user", user, true)
	if err != nil {
		return nil
	}

	if user.Admin {
		return user
	}

	return nil
}

func (c Admin) Index() revel.Result {
	if user := c.connected(); user != nil {
		return c.Render(user)
	}

	return c.Redirect(routes.App.Index())
}
