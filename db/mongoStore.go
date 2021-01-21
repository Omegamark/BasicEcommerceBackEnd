package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoStore struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
}

func (d *MongoStore) Ping() error {
	err := d.client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Errorf("Error connection to database: %v", err)
	}
	return err
}

func (d *MongoStore) GetStuff() ([]map[string]interface{}, error) {
	collection := d.client.Database(d.databaseName).Collection(d.collectionName)

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to find in db")
	}

	matches := []map[string]interface{}{}
	for cur.Next(context.Background()) {
		elem := map[string]interface{}{}
		err := cur.Decode(&elem)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal from db")
		}

		matches = append(matches, elem)
	}

	_, err = json.Marshal(matches)
	if err != nil {
		fmt.Println("failed to marshal", err)
	}
	return matches, nil
}

func (d *MongoStore) InsertStuff(thing map[string]interface{}) error {
	collection := d.client.Database(d.databaseName).Collection(d.collectionName)
	_, err := collection.InsertOne(context.Background(), thing)
	if err != nil {
		return err
	}

	return nil
}
