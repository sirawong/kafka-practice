package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"consumer/config"
	"consumer/domain"
)

type database struct {
	Client *mongo.Client
	DB     *mongo.Database
	Coll   *mongo.Collection
}

type DatabaseRepo interface{
	Save(ctx context.Context, ent *domain.Message) error
}

func New(dbConn *mongo.Client, appConfig *config.Config) DatabaseRepo {
	repo := &database{
		Client: dbConn,
		DB: dbConn.Database(appConfig.MongoDBCollection),
	}

	repo.Coll = repo.DB.Collection(appConfig.MongoDBCollection)

	return repo
}


func (d database) Save(ctx context.Context, ent *domain.Message) error {
	_, err := d.Coll.InsertOne(ctx, ent)
	if err != nil {
		return err
	}

	return nil
}