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
	clientOnce.Do(func() {
		// Initialize the MongoDB client here.
		clientOptions := options.Client().ApplyURI(getDatabaseConnectionString())
		fmt.Println(getDatabaseConnectionString())
		fmt.Println("test")

		c, err := mongo.Connect(context.Background(), clientOptions)

		if err != nil {
			panic(err)
		}
		client = c
	})

	return client, nil
}

// GetDatabase returns a MongoDB database from the client.
func GetDatabase(name string) (*mongo.Database, error) {
	client, _ := GetDBClient()
	return client.Database(name), nil
}

// GetTestCollection returns a MongoDB collection from the client's default database.
func GetTestCollection() (*mongo.Collection, error) {
	database, _ := GetDatabase(getDatabaseName())
	return database.Collection(getCollectionName()), nil
}

func getDatabaseConnectionString() string {
	return os.Getenv("DB_HOST_TEST")
}

func getDatabaseName() string {
	return os.Getenv("DB_NAME_TEST")
}

func getCollectionName() string {
	return os.Getenv("DB_COLLECTION_TEST")
}
