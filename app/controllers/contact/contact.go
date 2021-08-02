package contact

import (
	"github.com/revel/revel"

	"natka/app/controllers/utils"
	"natka/app/routes"
)

type Contact struct {
	*revel.Controller
}

func (c Contact) Index() revel.Result {
	_ = utils.IsConnected(c.Session)
	return c.Render()
}

// TODO: Fix missing textarea
func (c Contact) SendMessage(name string, mail string, text string, aggrement bool) revel.Result {
	c.Log.Infof("%v %v %v %v", name, mail, text, aggrement)

	return c.Redirect(routes.Contact.Index())
}
