package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"natka/app/models"
)

const articlesCollection = "articles"

type ArticleWithThumbnail struct {
	Article   models.Article
	Thumbnail models.Image
}

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

func GetArticleWithThumbnail(id *primitive.ObjectID) (*ArticleWithThumbnail, error) {
	rawArticle, err := GetArticle(id)
	if err != nil {
		return nil, err
	}

	thumbnail, err := GetImage(rawArticle.Thumbnail)
	if err != nil {
		return nil, err
	}

	return &ArticleWithThumbnail{
		Article:   *rawArticle,
		Thumbnail: thumbnail,
	}, nil
}

func GetArticles() ([]ArticleWithThumbnail, error) {
	var rawArticles []models.Article
	err := getAll(articlesCollection, bson.D{}, &options.FindOptions{Sort: bson.D{{Key: "_id", Value: -1}}}, &rawArticles)
	if err != nil {
		return nil, err
	}

	var articles []ArticleWithThumbnail
	for _, article := range rawArticles {
		thumbnail, err := GetImage(article.Thumbnail)
		if err != nil {
			return nil, err
		}
		articles = append(articles, ArticleWithThumbnail{
			Article:   article,
			Thumbnail: thumbnail,
		})
	}

	return articles, nil
}

func GetLastArticles() ([]ArticleWithThumbnail, error) {
	var rawArticles []models.Article
	limit := int64(3)
	err := getAll(articlesCollection, bson.D{}, &options.FindOptions{
		Sort:  bson.D{{Key: "_id", Value: -1}},
		Limit: &limit,
	}, &rawArticles)
	if err != nil {
		return nil, err
	}
	var articles []ArticleWithThumbnail
	for _, article := range rawArticles {
		thumbnail, err := GetImage(article.Thumbnail)
		if err != nil {
			return nil, err
		}
		articles = append(articles, ArticleWithThumbnail{
			Article:   article,
			Thumbnail: thumbnail,
		})
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
