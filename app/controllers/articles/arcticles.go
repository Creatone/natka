package articles

import (
	"github.com/revel/revel"

	"natka/app/db"
	"natka/app/models"
	"natka/app/routes"
)

type Articles struct {
	*revel.Controller
}

func (c *Articles) Articles() revel.Result {
	articles, err := db.GetArticles()
	if err != nil {
		return c.RenderError(err)
	}

	return c.Render(articles)
}

func (c *Articles) Add() revel.Result {
	return c.Render()
}

func (c *Articles) Insert(name string, description string) revel.Result {
	article := models.Article{
		Name:        name,
		Description: description,
	}

	_, err := db.InsertArticle(article)
	if err != nil {
		return c.RenderError(err)
	}

	return c.Redirect(routes.Articles.Add())
}
