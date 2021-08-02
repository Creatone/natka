package shop

import (
	"github.com/revel/revel"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"natka/app/controllers/utils"
	"natka/app/db"
	"natka/app/models"
	"natka/app/routes"
)

func (c *Shop) ShowBasket() revel.Result {
	_ = utils.IsConnected(c.Session)
	var basket models.Basket
	_, _ = c.Session.GetInto("basket", &basket, true)

	return c.Render(basket)
}

func (c *Shop) AddToBasket(id string) revel.Result {
	var basket models.Basket
	_, err := c.Session.GetInto("basket", &basket, true)
	if err != nil {
		basket = *models.NewBasket()
	}

	dietID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.Flash.Error(err.Error())
	}

	diet, err := db.GetDiet(&dietID)
	if err != nil {
		c.Log.Errorf("couldn't obtain %v: %w", dietID, err)
		return c.Redirect(routes.Shop.Index())
	}

	basket.Diets[id] = *diet

	err = c.Session.Set("basket", basket)
	if err != nil {
		c.Flash.Error(err.Error())
	}

	c.Flash.Success(c.Message("basket.diet.add"))

	return c.Redirect(routes.Shop.Index())
}

func (c *Shop) DeleteFromBasket(id string) revel.Result {
	var basket models.Basket
	_, err := c.Session.GetInto("basket", &basket, true)
	if err != nil {
		return c.Redirect(routes.Shop.Index())
	}

	basket.Delete(id)

	c.Flash.Success(c.Message("basket.diet.remove"))

	return c.Redirect(routes.Shop.ShowBasket())
}

func (c *Shop) FinalizeOptions() revel.Result {
	// TODO: Option without login.

	return c.Render()
}

func (c *Shop) Finalize() revel.Result {
	var basket models.Basket
	_, err := c.Session.GetInto("basket", &basket, true)
	if err != nil {
		c.Flash.Error(err.Error())
	}

	if len(basket.Diets) > 0 {
		if sessionUser := utils.IsConnected(c.Session); sessionUser != nil {
			if sessionUser.Diets == nil {
				sessionUser.Diets = make(map[string]struct{})
			}
			for k, _ := range basket.Diets {
				sessionUser.Diets[k] = struct{}{}
			}
			err := db.EditUser(*sessionUser)
			if err != nil {
				c.Flash.Error(err.Error())
				return c.Redirect(routes.Shop.ShowBasket())
			}
			err = c.Session.Set("user", sessionUser)
			if err != nil {
				c.Flash.Error(err.Error())
				return c.Redirect(routes.Shop.ShowBasket())
			}
			c.Flash.Success(c.Message("basket.order"))
			return c.Redirect(routes.Profile.Index())
		} else {
			return c.Redirect(routes.Shop.ShowBasket())
		}
	} else {
		return c.Redirect(routes.Shop.Index())
	}
}
