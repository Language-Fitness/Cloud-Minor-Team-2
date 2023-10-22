package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
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
		clientOptions := options.Client().ApplyURI(getDatabaseConnectionString())
		c, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			panic(err)
		}
		client = c
	})

	return client, nil
}

func getDatabaseConnectionString() string {
	return os.Getenv("DB_HOST")
}
