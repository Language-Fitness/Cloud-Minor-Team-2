package Repository

import (
	"Saga_orchestrator/src/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var mongoClient *mongo.Client
var sagaCollection *mongo.Collection

type SagaInstance = models.SagaInstance

// InitMongoDB Initialize MongoDB connection and collection
func InitMongoDB() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	mongoClient = client
	sagaCollection = mongoClient.Database("Saga-Test").Collection("sagas")
}

// CreateSagaInMongoDB Create a new saga and save it in MongoDB
func CreateSagaInMongoDB(saga *models.SagaInstance) error {
	_, err := sagaCollection.InsertOne(context.Background(), saga)
	return err
}

// GetSagaFromMongoDB Retrieve a saga from MongoDB by ID
func GetSagaFromMongoDB(sagaID string) (*models.SagaInstance, error) {
	filter := bson.M{"id": sagaID}
	var result models.SagaInstance
	err := sagaCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateSagaInMongoDB Update a saga in MongoDB
func UpdateSagaInMongoDB(saga *models.SagaInstance) error {
	filter := bson.M{"id": saga.ID}
	update := bson.M{
		"$set": bson.M{
			"currentStep": saga.CurrentStep,
			"state":       saga.State,
			"steps":       saga.Steps,
		},
	}
	_, err := sagaCollection.UpdateOne(context.Background(), filter, update)
	return err
}
