package app

import (
	"github.com/revel/revel"

	"natka/app/controllers/utils"
	"natka/app/db"
	"natka/app/models/instagram"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	siteTitle := "Dietetyk Natalia Danio"

	user := utils.IsConnected(c.Session)

	diets, _ := db.GetDiets()

	instaPosts, err := instagram.GetPosts()
	if err != nil {
		return c.RenderError(err)
	}

	return c.Render(siteTitle, user, diets, instaPosts)
}

func (c App) About() revel.Result {
	return c.Render()
}

func (c App) Calculator() revel.Result {
	return c.Render()
}

func (c App) Contact() revel.Result {
	return c.Render()
}
