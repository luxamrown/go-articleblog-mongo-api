package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mohamadelabror.me/goapiblog/model"
)

type ArticleRepo interface {
	CreateArticle(title, desc, author, date, article_id string) error
	GetAllArticle() ([]model.ArticleHead, error)
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

func (a *articleRepoImpl) GetAllArticle() ([]model.ArticleHead, error) {
	coll := a.articleDb.Database("mohamadelabror-blog").Collection("Blog")
	// findOptions := options.Find()
	var results []model.ArticleHead

	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []model.ArticleHead{}, err
	}

	for cursor.Next(context.TODO()) {
		var elem model.ArticleHead
		err := cursor.Decode(&elem)
		if err != nil {
			return []model.ArticleHead{}, err
		}
		results = append(results, elem)
	}
	return results, nil
}

func (a *articleRepoImpl) CreateArticle(title, desc, author, date, article_id string) error {
	coll := a.articleDb.Database("mohamadelabror-blog").Collection("Blog")
	doc := bson.M{"title": title, "desc": desc, "author": author, "date": date, "article_id": article_id}

	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}
	return nil

}

func NewArticleRepo(articleDb *mongo.Client) ArticleRepo {
	articleRepo := articleRepoImpl{
		articleDb: articleDb,
	}
	return &articleRepo
}
