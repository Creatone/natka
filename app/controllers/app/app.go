package app

import (
	"github.com/revel/revel"

	"natka/app/controllers/utils"
	"natka/app/db"
	"natka/app/routes"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	siteTitle := "Dietetyk Natalia Danio"

	if user := utils.IsConnected(c.Session); user != nil {
		siteTitle += " " + user.Mail
		diets, _ := db.GetDiets()
		return c.Render(user, diets)
	}

	return c.Redirect(routes.Login.Index())
}

func (c App) Calculator() revel.Result {
	return c.Render()
}
