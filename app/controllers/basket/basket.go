package basket

import "github.com/revel/revel"

type Basket struct {
	*revel.Controller
}

func (c *Basket) Basket() revel.Result {
	return c.Render()
}
