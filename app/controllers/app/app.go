package app

import (
	"github.com/revel/revel"

	"natka/app/db"
	"natka/app/models"
	"natka/app/routes"
)

type App struct {
	*revel.Controller
}

func (c App) connected() *models.User {
	user := &models.User{}
	_, err := c.Session.GetInto("user", user, true)
	if err != nil {
		return nil
	}

	return user
}

func (c App) Index() revel.Result {
	siteTitle := "Dietetyk Natalia Danio"

	if user := c.connected(); user != nil {
		siteTitle += " " + user.Mail
		diets, _ := db.GetDiets()
		return c.Render(user, diets)
	}

	return c.Redirect(routes.Login.Index())
}
