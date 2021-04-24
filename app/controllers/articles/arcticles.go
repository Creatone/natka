package articles

import "github.com/revel/revel"

type Articles struct {
	*revel.Controller
}

func (c *Articles) Articles() revel.Result {
	return c.Render()
}
