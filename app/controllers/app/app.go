package app

import (
	"encoding/base64"

	"github.com/revel/revel"

	"natka/app/controllers/utils"
	"natka/app/db"
	"natka/app/models"
	"natka/app/models/instagram"
	"natka/app/routes"
)

const (
	carousel = "carousel"
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

	diets, _ := db.GetDiets()

	carousel, _ := db.GetImagesByType(carousel)

	return c.Render(siteTitle, user, instaPosts, diets, carousel)
}

func (c App) AddCarousel() revel.Result {
	if sessionUser := utils.IsConnected(c.Session); sessionUser != nil && sessionUser.Admin {
		return c.Render()
	}
	return c.Redirect(routes.App.Index())
}

func (c App) InsertCarousel(carouselimage []byte) revel.Result {
	if sessionUser := utils.IsConnected(c.Session); sessionUser != nil && sessionUser.Admin {
		image := models.Image{
			Data: base64.StdEncoding.EncodeToString(carouselimage),
			Type: carousel,
		}

		_, err := db.InsertImage(image)
		if err != nil {
			c.Flash.Error(err.Error())
		}
	}

	return c.Redirect(routes.App.Index())
}

func (c App) EditCarousel(id string) revel.Result {
	if sessionUser := utils.IsConnected(c.Session); sessionUser != nil && sessionUser.Admin {
		return c.Render(id)
	}
	return c.Redirect(routes.App.Index())
}

func (c App) ApplyEditCarousel(carouselimage []byte, id string) revel.Result {
	if sessionUser := utils.IsConnected(c.Session); sessionUser != nil && sessionUser.Admin {
		image := models.Image{
			ID:   id,
			Data: base64.StdEncoding.EncodeToString(carouselimage),
			Type: carousel,
		}

		err := db.EditImage(image)
		if err != nil {
			c.Flash.Error(err.Error())
			return c.Redirect(routes.App.Index())
		}
	}

	return c.Redirect(routes.App.Index())
}

func (c App) DeleteCarousel(id string) revel.Result {
	if sessionUser := utils.IsConnected(c.Session); sessionUser != nil && sessionUser.Admin {
		err := db.DeleteImage(id)
		if err != nil {
			c.Flash.Error(err.Error())
		}
	}
	return c.Redirect(routes.App.Index())
}

func (c App) About() revel.Result {
	_ = utils.IsConnected(c.Session)
	return c.Render()
}

func (c App) Calculator() revel.Result {
	_ = utils.IsConnected(c.Session)
	return c.Render()
}
