package app

import (
	"github.com/revel/revel"

	"natka/app/controllers/utils"
	"natka/app/db"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	siteTitle := "Dietetyk Natalia Danio"

	user := utils.IsConnected(c.Session)

	siteTitle += " " + user.Mail
	diets, _ := db.GetDiets()
	return c.Render(user, diets)
}

func (c App) Calculator() revel.Result {
	return c.Render()
}
