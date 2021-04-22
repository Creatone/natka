package profile

import (
	"bytes"
	"encoding/base64"
	img "image"
	_ "image/jpeg"

	"github.com/revel/revel"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"natka/app/controllers/utils"
	"natka/app/db"
	"natka/app/models"
	"natka/app/routes"
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
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
		c.Validation.Required(avatar)
		c.Validation.MinSize(avatar, 2*KB)
		c.Validation.MaxSize(avatar, 4*MB)
		rawImage, format, err := img.DecodeConfig(bytes.NewReader(avatar))
		c.Validation.Required(err == nil).Key("avatar").Message("Incorrect file format")
		c.Validation.Min(rawImage.Height, 600).Message("Minimum avatar size is 600x600")
		c.Validation.Min(rawImage.Width, 600).Message("Minimum avatar size is 600x600")
		c.Validation.Required(format == "jpeg").Key("avatar").Message("JPEG format is required")
		if c.Validation.HasErrors() {
			c.Validation.Keep()
			c.FlashParams()
			return c.Redirect(routes.Profile.Edit())
		}
		// TODO: Shrink image.
		image := models.Image{Data: avatar}
		if sessionUser.Avatar == nil {
			id, err := db.InsertImage(image)
			var another = id.(primitive.ObjectID)
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
		err = db.EditUser(*sessionUser)
		if err != nil {
			return c.RenderError(err)
		}
	}

	return c.Redirect(routes.Profile.Index())
}
