package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mohamadelabror.me/goapiblog/model"
)

type ArticleRepo interface {
	GetArticle(id string) (model.Article, error)
}

type articleRepoImpl struct {
	articleDb *mongo.Client
}

func (a *articleRepoImpl) GetArticle(id string) (model.Article, error) {
	var selectedArticle model.Article
	coll := a.articleDb.Database("mohamadelabror-blog").Collection("Blog")

	err := coll.FindOne(context.TODO(), bson.M{"article_id": id}).Decode(&selectedArticle)
	if err != nil {
		return model.Article{}, err
	}

	return selectedArticle, nil
}

func NewArticleRepo(articleDb *mongo.Client) ArticleRepo {
	articleRepo := articleRepoImpl{
		articleDb: articleDb,
	}
	return &articleRepo
}
