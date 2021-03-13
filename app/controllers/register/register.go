package register

import (
	"fmt"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"

	"natka/app/db"
	"natka/app/models"
	"natka/app/routes"
)

type Register struct {
	*revel.Controller
}

func (c Register) Index(mail string) revel.Result {
	return c.Render(mail)
}

func (c Register) Register(mail, name, password string) revel.Result {
	c.Session.Get("user")

	user := models.User{
		Name: name,
		Mail: mail,
	}

	var err error
	user.Password, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Errorf("cannot encrypt %q password: %v", user.Name, err)
	} else {
		err = db.InsertUser(user)
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
