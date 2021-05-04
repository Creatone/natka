package db

import (
	"go.mongodb.org/mongo-driver/bson"

	"natka/app/models"
)

const articlesCollection = "articles"

func InsertArticle(article models.Article) (interface{}, error) {
	return insert(articlesCollection, article)
}

func GetArticles() ([]models.Article, error) {
	var articles []models.Article

	err := getAll(articlesCollection, bson.D{}, &articles)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
