package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"natka/app/models"
)

const articlesCollection = "articles"

func InsertArticle(article models.Article) (interface{}, error) {
	return insert(articlesCollection, article)
}

func GetArticle(id *primitive.ObjectID) (*models.Article, error) {
	article := models.Article{}

	err := get(articlesCollection, bson.D{{"_id", id}}, &article)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func GetArticles() ([]models.Article, error) {
	var articles []models.Article

	err := getAll(articlesCollection, bson.D{}, &articles)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func EditArticle(article models.Article) error {
	// TODO: Move bson-key thing to func.
	objectID, err := primitive.ObjectIDFromHex(article.ID)
	if err != nil {
		return err
	}
	return edit(articlesCollection, bson.D{{"_id", objectID}},
		bson.D{{"$set", bson.D{
			{"name", article.Name},
			{"description", article.Description},
			{"text", article.Text},
		}}})
}
