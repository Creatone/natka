package app

import (
	"fmt"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"

	"natka/app/db"
	"natka/app/models"
	"natka/app/routes"
)

func (c App) Register(mail string) revel.Result {
	if !c.Validation.Email(mail).Ok {
		c.FlashParams()

		return c.Redirect(routes.App.Index())
	}

	// Check if email is in DB
	isMail, err := db.CheckMail(mail)
	if isMail {
		return c.Redirect(routes.App.EnterPassword(mail))
	}
	if err != nil {
		return c.RenderError(fmt.Errorf("error checking email: %v", err))
	}
	// If not -> go to registry

	return c.Render(mail)
}

func (c App) ChangeName(name string) revel.Result {
	return c.Redirect(routes.App.Index())
}

func (c App) SaveUser(user models.User, password string) revel.Result {
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Register(user.Mail))
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

	return c.Redirect(routes.App.Index())
}
