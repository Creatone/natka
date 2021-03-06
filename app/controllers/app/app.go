package app

import (
	"github.com/revel/revel"
	"natka/app/db"
	"natka/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) connected() *models.User {

	if c.ViewArgs["fulluser"] != nil {
		return c.ViewArgs["fulluser"].(*models.User)
	}

	if mail, ok := c.Session["mail"]; ok {
		return c.getUser(mail.(string))
	}

	return nil
}

func (c App) getUser(mail string) *models.User {
	user := &models.User{}
	_, err := c.Session.GetInto("fulluser", user, false)
	if user.Mail == mail {
		return user
	}

	user, err = db.GetUser(mail)
	if err != nil {
		c.Log.Errorf("failed to authorize user: %v", err)
	}

	return user
}

func (c App) Index() revel.Result {
	siteTitle := "Dietetyk Natalia Danio"

	if user := c.connected(); user != nil {
		siteTitle += " " + user.Mail
	}

	return c.Render(siteTitle)
}
