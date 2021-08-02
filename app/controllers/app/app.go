package app

import (
	"github.com/revel/revel"

	"natka/app/controllers/utils"
	"natka/app/models/instagram"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	siteTitle := "Dietetyk Natalia Danio"

	user := utils.IsConnected(c.Session)

	instaPosts, err := instagram.GetPosts()
	if err != nil {
		return c.RenderError(err)
	}

	return c.Render(siteTitle, user, instaPosts)
}

func (c App) About() revel.Result {
	_ = utils.IsConnected(c.Session)
	return c.Render()
}

func (c App) Calculator() revel.Result {
	_ = utils.IsConnected(c.Session)
	return c.Render()
}
