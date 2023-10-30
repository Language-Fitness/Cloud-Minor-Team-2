package graph

import (
	"Module/graph/model"
	"Module/internal/repository"
	database "Module/test/internal/helpers"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
	"time"
)

func TestCreateModule(t *testing.T) {
	// Get the MongoDB collection for testing.
	fmt.Println("Starting")

	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	// Create a new ModuleRepository using the test collection.
	repo := repository.NewModuleRepository(collection)

	// Define your test data based on the Module struct.
	desc := "Test description"
	category := "Test category"
	user := "Test user"
	key := "test-key"
	createdAt := "2023-10-27"
	difficulty := 100
	boolean := false

	newModule := &model.Module{
		ID:          "123",
		Name:        "Test Module",
		Description: &desc,
		Difficulty:  &difficulty,
		Category:    &category,
		MadeBy:      &user,
		Private:     &boolean,
		Key:         &key,
		CreatedAt:   &createdAt,
		UpdatedAt:   &createdAt,
		SoftDeleted: &boolean,
	}

	// Call the method you want to test.
	_, err = repo.CreateModule(newModule)

	// Assert the result and error as needed.
	if err != nil {
		t.Errorf("Error creating module: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()
	// Attempt to fetch the module from MongoDB.
	filter := bson.M{"id": "123"}
	var databaseModule model.Module

	err = collection.FindOne(ctx, filter).Decode(&databaseModule)

	if err != nil {
		t.Errorf("Error fetching module from MongoDB: %v", err)
	}

	if !reflect.DeepEqual(newModule, &databaseModule) {
		t.Errorf("Retrieved module does not match the expected module")
	}
}
