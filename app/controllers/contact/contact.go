package contact

import (
	"natka/app/routes"

	"github.com/revel/revel"
)

type Contact struct {
	*revel.Controller
}

func (c Contact) Index() revel.Result {
	return c.Render()
}

// TODO: Fix missing textarea
func (c Contact) SendMessage(name string, mail string, text string, aggrement bool) revel.Result {
	c.Log.Infof("%v %v %v %v", name, mail, text, aggrement)

	return c.Redirect(routes.Contact.Index())
}
