package profile

import (
	"encoding/base64"

	"github.com/revel/revel"

	"go.mongodb.org/mongo-driver/bson/primitive"

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
		var image models.Image
		if user.Avatar != nil {
			var err error
			image, err = db.GetImage(user.Avatar)
			if err != nil {
				return c.RenderError(err)
			}
		}
		avatar := base64.StdEncoding.EncodeToString(image.Data)
		return c.Render(user, diets, avatar)
	}

	return c.Redirect(routes.App.Index())
}

func (c Profile) Edit() revel.Result {
	if user := utils.IsConnected(c.Session); user != nil {
		var image models.Image
		if user.Avatar != nil {
			var err error
			image, err = db.GetImage(user.Avatar)
			if err != nil {
				return c.RenderError(err)
			}
		}
		avatar := base64.StdEncoding.EncodeToString(image.Data)
		return c.Render(user, avatar)
	}

	return c.Redirect(routes.App.Index())
}

func (c Profile) ApplyEdit(user models.User, avatar []byte) revel.Result {
	if sessionUser := utils.IsConnected(c.Session); sessionUser != nil {
		sessionUser.Name = user.Name
		if len(avatar) != 0 {
			image := models.Image{Data: avatar}
			if sessionUser.Avatar == nil {
				id, err := db.InsertImage(image)
				var another primitive.ObjectID = id.(primitive.ObjectID)
				sessionUser.Avatar = &another
				if err != nil {
					return c.RenderError(err)
				}
			} else {
				image.ID = sessionUser.Avatar
				err := db.EditImage(image)
				if err != nil {
					return c.RenderError(err)
				}
			}
		}
		err := db.EditUser(*sessionUser)
		if err != nil {
			return c.RenderError(err)
		}
	}

	return c.Redirect(routes.Profile.Index())
}
