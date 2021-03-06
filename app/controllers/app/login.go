package app

import (
	"fmt"
	"github.com/revel/revel"
	"natka/app/db"
	"natka/app/routes"
)

func (c App) EnterPassword(mail string) revel.Result {
	return c.Render(mail)
}

func (c App) Login(mail, password string) revel.Result {
	pass, err := db.Login(mail, password)
	if err != nil {
		return c.RenderError(err)
	}
	if pass {
		c.Session.SetDefaultExpiration()
		err := c.Session.Set("mail", mail)
		if err != nil {
			return c.RenderError(err)
		}
		return c.Redirect(routes.App.Index())
	}

	return c.RenderError(fmt.Errorf("cannot login"))
}

func (c App) Logout() revel.Result {
	c.Session.Del("fulluser")
	c.Session.Del("mail")

	return c.Redirect(routes.App.Index())
}
