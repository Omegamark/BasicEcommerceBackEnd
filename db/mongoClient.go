package db

import (
	"context"
	configuration "reactApp/mongoClient/config"
	handlers "reactApp/mongoClient/handler"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitClient(ctx context.Context) (*mongo.Client, error) {
	mClient, err := mongo.Connect(ctx, options.Client().ApplyURI(configuration.MongoURI))
	if err != nil {
		return nil, err
	}

	err = mClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return mClient, nil
}

func InitDBs(mClient *mongo.Client, conf configuration.Config) handlers.MongoDataStore {
	ms := &MongoStore{
		client:         mClient,
		databaseName:   conf.DBName,
		collectionName: conf.CollName,
	}

	return ms
}
