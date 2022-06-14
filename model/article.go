package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title,omitempty"`
	Desc      string             `bson:"desc,omitempty"`
	Author    string             `bson:"author,omitempty"`
	ArticleId string             `bson:"article_id,ompitempty"`
	Date      string             `bson:"date,omitempty"`
}

type ArticleHead struct {
	Title     string `bson:"title,omitempty"`
	Author    string `bson:"author,omitempty"`
	ArticleId string `bson:"article_id,ompitempty"`
	Date      string `bson:"date,omitempty"`
}
