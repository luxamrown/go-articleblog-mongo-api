package manager

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Infra interface {
	MongoDb() *mongo.Client
}

type infra struct {
	db *mongo.Client
}

func (i *infra) MongoDb() *mongo.Client {
	return i.db
}

func NewInfra(dataSourceName string) Infra {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dataSourceName))
	if err != nil {
		panic(err)
	}

	return &infra{
		db: client,
	}
}
