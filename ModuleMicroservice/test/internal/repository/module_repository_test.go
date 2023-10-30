package graph

import (
	"Module/graph/model"
	"Module/internal/repository"
	database "Module/test/internal/helpers"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestCreateModule(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
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

func TestUpdateModule(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new ModuleRepository using the test collection.
	repo := repository.NewModuleRepository(collection)

	// Define your test data for an existing module.
	existingModule := &model.Module{
		ID:   "123",
		Name: "Test Module",
		// Initialize other fields as needed for the existing module.
	}

	// Insert the existing module into MongoDB.
	_, err = repo.CreateModule(existingModule)
	if err != nil {
		t.Errorf("Error creating the existing module: %v", err)
	}

	// Define the updates you want to apply to the module.
	updatedModuleInput := model.ModuleInput{
		Name: "Updated Test Module",
		// Define other fields you want to update.
	}

	// Call the method you want to test.
	updatedModule, err := repo.UpdateModule("123", updatedModuleInput)

	// Assert the result and error as needed.
	if err != nil {
		t.Errorf("Error updating module: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()
	// Attempt to fetch the updated module from MongoDB.
	filter := bson.M{"id": "123"}
	var databaseModule model.Module

	err = collection.FindOne(ctx, filter).Decode(&databaseModule)

	if err != nil {
		t.Errorf("Error fetching updated module from MongoDB: %v", err)
	}

	// Assert that the updated module fields match the expected updates.
	if updatedModule.Name != updatedModuleInput.Name {
		t.Errorf("Updated module name does not match the expected value")
	}
	// Add similar assertions for other fields you updated.
}

func TestDeleteModuleByID(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new ModuleRepository using the test collection.
	repo := repository.NewModuleRepository(collection)

	// Define your test data for an existing module.
	existingModule := &model.Module{
		ID:   "123",
		Name: "Test Module",
		// Initialize other fields as needed for the existing module.
	}

	// Insert the existing module into MongoDB.
	_, err = repo.CreateModule(existingModule)
	if err != nil {
		t.Errorf("Error creating the existing module: %v", err)
	}

	// Call the method you want to test.
	err = repo.DeleteModuleByID("123")

	// Assert the error as needed.
	if err != nil {
		t.Errorf("Error deleting module: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()
	// Attempt to fetch the deleted module from MongoDB.
	filter := bson.M{"id": "123"}
	var databaseModule model.Module

	err = collection.FindOne(ctx, filter).Decode(&databaseModule)

	// Assert that the error is not nil, indicating the module was deleted.
	if err == nil {
		t.Errorf("Module was not deleted as expected")
	}
}

func TestGetModuleByID(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new ModuleRepository using the test collection.
	repo := repository.NewModuleRepository(collection)

	// Define your test data for an existing module.
	existingModule := &model.Module{
		ID:   "123",
		Name: "Test Module",
		// Initialize other fields as needed for the existing module.
	}

	// Insert the existing module into MongoDB.
	_, err = repo.CreateModule(existingModule)
	if err != nil {
		t.Errorf("Error creating the existing module: %v", err)
	}

	// Call the method you want to test.
	fetchedModule, err := repo.GetModuleByID("123")

	// Assert the result and error as needed.
	if err != nil {
		t.Errorf("Error fetching module: %v", err)
	}

	// Assert that the fetched module is not nil and has the expected ID.
	if fetchedModule == nil {
		t.Errorf("Fetched module is nil, expected a valid module")
	} else if fetchedModule.ID != "123" {
		t.Errorf("Fetched module has the wrong ID")
	}
}

func TestListModules(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new ModuleRepository using the test collection.
	repo := repository.NewModuleRepository(collection)

	// Define your test data for multiple modules.
	module1 := &model.Module{
		ID:   "1",
		Name: "Module 1",
		// Initialize other fields as needed for module1.
	}
	module2 := &model.Module{
		ID:   "2",
		Name: "Module 2",
		// Initialize other fields as needed for module2.
	}

	// Insert the test modules into MongoDB.
	_, err = repo.CreateModule(module1)
	if err != nil {
		t.Errorf("Error creating module 1: %v", err)
	}
	_, err = repo.CreateModule(module2)
	if err != nil {
		t.Errorf("Error creating module 2: %v", err)
	}

	// Call the method you want to test.
	modules, err := repo.ListModules()

	// Assert the result and error as needed.
	if err != nil {
		t.Errorf("Error listing modules: %v", err)
	}

	// Assert that the modules slice is not nil and contains the expected modules.
	if modules == nil {
		t.Errorf("List of modules is nil, expected a valid slice")
	} else if len(modules) != 2 {
		t.Errorf("Expected 2 modules, got %d", len(modules))
	}

	// Add specific assertions for each module in the list.
	if modules[0].ID != "1" || modules[1].ID != "2" {
		t.Errorf("Listed modules have incorrect IDs")
	}
}

func clearCollection(collection *mongo.Collection) error {
	_, err := collection.DeleteMany(context.TODO(), bson.D{})
	return err
}
