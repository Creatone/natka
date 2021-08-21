package diets

import (
	"github.com/revel/revel"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"natka/app/controllers/utils"
	"natka/app/db"
	"natka/app/models"
	"natka/app/routes"
)

type Diets struct {
	*revel.Controller
}

func (c *Diets) Add() revel.Result {
	return c.Render()
}

func (c *Diets) Insert(name string, description string) revel.Result {
	diet := models.Diet{
		Name:        name,
		Description: description,
	}

	_, err := db.InsertDiet(diet)
	if err != nil {
		return c.RenderError(err)
	}

	return c.Redirect(routes.Shop.Index())
}

func (c *Diets) Delete(id string) revel.Result {
	if sessionUser := utils.IsConnected(c.Session); sessionUser != nil && sessionUser.Admin {
		dietID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			// TODO: Format error.
			c.Flash.Error(err.Error())
			return c.Redirect(routes.Shop.Index())
		}

		err = db.DeleteDiet(&dietID)
		if err != nil {
			// TODO: Format error.
			c.Flash.Error(err.Error())
			return c.Redirect(routes.Shop.Index())
		}
		c.Flash.Success(c.Message("basket.diet.remove"))
		return c.Redirect(routes.Shop.Index())
	}

	return c.Redirect(routes.App.Index())
}

func (c *Diets) Edit(id string) revel.Result {
	dietID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// TODO: Format error.
		c.Flash.Error(err.Error())
		return c.Redirect(routes.Shop.Index())
	}

	diet, err := db.GetDiet(&dietID)
	if err != nil {
		// TODO: Format error.
		c.Flash.Error(err.Error())
		return c.Redirect(routes.Shop.Index())
	}

	return c.Render(diet)
}

func (c *Diets) EditApply(diet models.Diet) revel.Result {
	err := db.EditDiet(diet)
	if err != nil {
		// TODO: Format error.
		c.Flash.Error(err.Error())
		return c.Redirect(routes.Shop.Index())
	}

	// TODO: Success message.
	return c.Redirect(routes.Diets.Show(diet.ID))
}

func (c *Diets) Show(id string) revel.Result {
	dietID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.RenderError(err)
	}

	diet, err := db.GetDiet(&dietID)
	if err != nil {
		return c.RenderError(err)
	}

	return c.Render(diet)
}
