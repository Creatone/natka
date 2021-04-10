package admin

import (
	"github.com/revel/revel"

	"natka/app/controllers/utils"
	"natka/app/db"
	"natka/app/routes"
)

type Admin struct {
	*revel.Controller
}

func (c Admin) Index() revel.Result {
	if user := utils.IsConnected(c.Session); user != nil && user.Admin {
		diets, err := db.GetDiets()
		if err != nil {
			return c.RenderError(err)
		}
		return c.Render(user, diets)
	}

	return c.Redirect(routes.App.Index())
}
