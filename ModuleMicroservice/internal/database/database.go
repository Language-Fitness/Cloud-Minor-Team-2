package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var (
	client     *mongo.Client
	clientOnce sync.Once
)

// GetDBClient returns a singleton MongoDB client instance.
func GetDBClient() (*mongo.Client, error) {
	clientOnce.Do(func() {
		// Initialize the MongoDB client here.
		clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017")
		c, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			panic(err)
		}
		client = c
	})

	return client, nil
}
