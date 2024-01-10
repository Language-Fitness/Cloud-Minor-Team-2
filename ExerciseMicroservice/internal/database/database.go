package database

import (
	"context"
	"fmt"
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
	fmt.Println(getDatabaseConnectionString())

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

func GetDatabase() (*mongo.Database, error) {
	client, _ := GetDBClient()

	return client.Database(getDatabaseName()), nil
}

func GetCollection() (*mongo.Collection, error) {
	database, _ := GetDatabase()

	return database.Collection(getCollectionName()), nil
}

func getDatabaseConnectionString() string {
	return os.Getenv("DB_HOST")
}

func getDatabaseName() string {
	return os.Getenv("DB_NAME")
}

func getCollectionName() string {
	return os.Getenv("DB_COLLECTION")
}
