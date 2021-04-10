package login

import (
	"fmt"

	"github.com/revel/revel"

	"go.mongodb.org/mongo-driver/mongo"

	"golang.org/x/crypto/bcrypt"

	"natka/app/db"
	"natka/app/models"
	"natka/app/routes"
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

func isMailPresent(mail string) (bool, error) {
	_, err := db.GetUser(mail)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func login(mail, password string) (*models.User, error) {
	user, err := db.GetUser(mail)
	if err != nil {
		return nil, err
	}
	if user != nil {
		err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	return nil, nil
}

func (c Login) CheckMail(mail string) revel.Result {
	if !c.Validation.Email(mail).Ok {
		c.FlashParams()
		return c.Redirect(routes.App.Index())
	}

	isMail, err := isMailPresent(mail)
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
	user, err := login(mail, password)
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
