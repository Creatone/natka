package articles

import (
	"fmt"

	"github.com/revel/revel"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"natka/app/controllers/utils"
	"natka/app/db"
	"natka/app/models"
	"natka/app/routes"
)

type Articles struct {
	*revel.Controller
}

func (c *Articles) Articles() revel.Result {
	_ = utils.IsConnected(c.Session)

	articles, err := db.GetArticles()
	if err != nil {
		return c.RenderError(err)
	}

	return c.Render(articles)
}

func (c *Articles) Add() revel.Result {
	return c.Render()
}

func (c *Articles) Image() revel.Result {
	if user := utils.IsConnected(c.Session); user != nil && user.Admin {
		data := make(map[string]interface{})
		data["location"] = "/public/img/diet.jpg"
		return c.RenderJSON(data)
	}
	return c.RenderError(fmt.Errorf("Not an admin!"))
}

func (c *Articles) Insert(name string, description string, text string) revel.Result {
	article := models.Article{
		Name:        name,
		Description: description,
		Text:        text,
	}

	_, err := db.InsertArticle(article)
	if err != nil {
		return c.RenderError(err)
	}

	return c.Redirect(routes.Articles.Add())
}

func (c *Articles) Edit(id string) revel.Result {
	articleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.RenderError(err)
	}

	article, err := db.GetArticle(&articleID)
	if err != nil {
		return c.RenderError(err)
	}

	return c.Render(article)
}

func (c *Articles) EditApply(article models.Article) revel.Result {
	err := db.EditArticle(article)
	if err != nil {
		return c.RenderError(err)
	}

	return c.Redirect(routes.Articles.Articles())
}

func (c *Articles) Show(id string) revel.Result {
	articleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.RenderError(err)
	}

	article, err := db.GetArticle(&articleID)
	if err != nil {
		return c.RenderError(err)
	}

	return c.Render(article)
}
