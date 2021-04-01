package admin

import (
	"github.com/revel/revel"

	"natka/app/db"
	"natka/app/models"
	"natka/app/routes"
)

func (c Admin) IndexDiet() revel.Result {
	return c.Render()
}

func (c Admin) InsertDiet(name string, description string) revel.Result {
	diet := models.Diet{
		Name:        name,
		Description: description,
	}

	err := db.InsertDiet(diet)
	if err != nil {
		return c.RenderError(err)
	}

	return c.Redirect(routes.Admin.Index())
}
