package profile

import (
	"github.com/revel/revel"

	"natka/app/controllers/utils"
	"natka/app/db"
	"natka/app/models"
	"natka/app/routes"
)

type Profile struct {
	*revel.Controller
}

func (c Profile) Index() revel.Result {
	if user := utils.IsConnected(c.Session); user != nil {
		diets, _ := db.GetDiets()
		return c.Render(user, diets)
	}

	return c.Redirect(routes.App.Index())
}

func (c Profile) Edit() revel.Result {
	if user := utils.IsConnected(c.Session); user != nil {
		return c.Render(user)
	}

	return c.Redirect(routes.App.Index())
}

func (c Profile) ApplyEdit(user models.User) revel.Result {
	if sessionUser := utils.IsConnected(c.Session); sessionUser != nil {
		sessionUser.Name = user.Name
		err := db.EditUser(*sessionUser)
		if err != nil {
			return c.RenderError(err)
		}
	}

	return c.Redirect(routes.Profile.Index())
}
