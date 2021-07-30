package shop

import (
	"github.com/revel/revel"

	"natka/app/controllers/utils"
	"natka/app/db"
)

type Shop struct {
	*revel.Controller
}

func (c *Shop) Index() revel.Result {
	_ = utils.IsConnected(c.Session)

	diets, _ := db.GetDiets()

	return c.Render(diets)
}
