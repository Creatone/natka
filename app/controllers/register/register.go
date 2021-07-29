package register

import (
	"fmt"

	"github.com/revel/revel"
	"github.com/revel/revel/session"

	"golang.org/x/crypto/bcrypt"

	"natka/app/db"
	"natka/app/models"
	"natka/app/routes"
)

type Register struct {
	*revel.Controller
}

func (c Register) Index(mail string) revel.Result {
	sessionUser, err := c.Session.Get("user")
	if err != nil {
		if err.Error() != session.SESSION_VALUE_NOT_FOUND.Error() {
			c.Flash.Error(err.Error())
			return c.Redirect(routes.App.Index())
		}
	}

	if sessionUser != nil {
		return c.Redirect(routes.App.Index())
	}

	return c.Render(mail)
}

func (c Register) Register(mail, name, password string) revel.Result {
	sessionUser, err := c.Session.Get("user")
	if err != nil {
		if err.Error() != session.SESSION_VALUE_NOT_FOUND.Error() {
			c.Flash.Error(err.Error())
			return c.Redirect(routes.App.Index())
		}
	}

	if sessionUser != nil {
		return c.Redirect(routes.App.Index())
	}

	user := models.User{
		Name: name,
		Mail: mail,
	}

	user.Password, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Errorf("cannot encrypt %q password: %v", user.Name, err)
	} else {
		_, err = db.InsertUser(user)
		if err != nil {
			return c.RenderError(fmt.Errorf("cannot insert user: %v", err))
		}
	}

	// Then login user.
	c.Session.SetDefaultExpiration()
	err = c.Session.Set("user", user)
	if err != nil {
		return c.RenderError(err)
	}

	return c.Redirect(routes.App.Index())
}
