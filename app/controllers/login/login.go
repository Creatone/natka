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

const (
	wrongPasswordErrorKey = "login.wrong.password"
	noEmailErrorKey       = "login.no.email"
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

func login(mail, password string, noUserErrorString string, wrongPasswordErrorString string) (*models.User, error) {
	user, err := db.GetUser(mail)
	if err != nil {
		return nil, fmt.Errorf(noUserErrorString)
	}
	if user != nil {
		err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
		if err != nil {
			return nil, fmt.Errorf(wrongPasswordErrorString)
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
	user, err := login(mail, password, c.Message(noEmailErrorKey), c.Message(wrongPasswordErrorKey))
	if err != nil {
		c.Validation.Error(err.Error())
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Login.EnterPassword(mail))
	}
	if user != nil {
		c.Session.SetDefaultExpiration()
		err := c.Session.Set(sessionKey, user)
		if err != nil {
			c.Validation.Error(err.Error())
			c.Validation.Keep()
			return c.Redirect(routes.Login.Index())
		}
		return c.Redirect(routes.App.Index())
	}
	return c.Redirect(routes.Login.Index())
}

func (c Login) Logout() revel.Result {
	c.Session.Del(sessionKey)

	return c.Redirect(routes.App.Index())
}
