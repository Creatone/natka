package login

import (
	"fmt"
	"github.com/revel/revel"
	"natka/app/routes"

	"natka/app/db"
)

var sessionKey = "user"

type Login struct {
	*revel.Controller
}

func (c Login) Index() revel.Result {
	return c.Render()
}

func (c Login) EnterPassword(mail string) revel.Result {
	return c.Render(mail)
}

func (c Login) CheckMail(mail string) revel.Result {
	if !c.Validation.Email(mail).Ok {
		c.FlashParams()
		return c.Redirect(routes.App.Index())
	}

	isMail, err := db.CheckMail(mail)
	if err != nil {
		return c.Redirect(routes.App.Index())
	}

	if isMail {
		return c.Redirect(routes.Login.EnterPassword(mail))
	} else {
		return c.Redirect(routes.Register.Index(mail))
	}
}

func (c Login) Login(mail, password string) revel.Result {
	user, err := db.Login(mail, password)
	if err != nil {
		return c.RenderError(err)
	}
	if user != nil {
		c.Session.SetDefaultExpiration()
		err := c.Session.Set(sessionKey, user)
		if err != nil {
			return c.RenderError(err)
		}
		return c.Redirect(routes.App.Index())
	}

	return c.RenderError(fmt.Errorf("cannot login"))
}

func (c Login) Logout() revel.Result {
	c.Session.Del(sessionKey)

	return c.Redirect(routes.App.Index())
}
