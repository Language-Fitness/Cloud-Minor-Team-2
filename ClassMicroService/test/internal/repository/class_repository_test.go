package graph

import (
	"context"
	"example/graph/model"
	"example/internal/repository"
	database "example/test/internal/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestCreateClass(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new ClassRepository using the test collection.
	repo := repository.NewClassRepository(collection)

	// Define your test data based on the Class struct.
	description := "Test description."
	difficulty := 1
	timestamp := time.Now().String()
	softDeleted := false

	newClass := &model.Class{
		ID:          "123",
		ModuleID:    "Test module-id",
		Name:        "Test Class",
		Description: description,
		Difficulty:  difficulty,
		CreatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	}

	// Call the method you want to test.
	_, err = repo.CreateClass(newClass)

	// Assert the result and error as needed.
	if err != nil {
		t.Errorf("Error creating class: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()
	// Attempt to fetch the class from MongoDB.
	filter := bson.M{"id": "123"}
	var databaseClass model.Class

	err = collection.FindOne(ctx, filter).Decode(&databaseClass)

	if err != nil {
		t.Errorf("Error fetching class from MongoDB: %v", err)
	}

	if !reflect.DeepEqual(newClass, &databaseClass) {
		t.Errorf("Retrieved class does not match the expected class")
	}
}

func TestUpdateClass(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new ClassRepository using the test collection.
	repo := repository.NewClassRepository(collection)

	// Define your test data for an existing class.
	existingClass := &model.Class{
		ID:       "123",
		ModuleID: "module-id",
		Name:     "Test Class",
		// Initialize other fields as needed for the existing class.
	}

	// Insert the existing class into MongoDB.
	_, err = repo.CreateClass(existingClass)
	if err != nil {
		t.Errorf("Error creating the existing class: %v", err)
	}

	// Define the updates you want to apply to the class.
	updatedClassInput := model.Class{
		ID:       "123",
		ModuleID: "module-id",
		Name:     "Updated Test Class",
		// Define other fields you want to update.
	}

	// Call the method you want to test.
	updatedClass, err := repo.UpdateClass("123", updatedClassInput)

	// Assert the result and error as needed.
	if err != nil {
		t.Errorf("Error updating class: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()
	// Attempt to fetch the updated class from MongoDB.
	filter := bson.M{"id": "123"}
	var databaseClass model.Class

	err = collection.FindOne(ctx, filter).Decode(&databaseClass)

	if err != nil {
		t.Errorf("Error fetching updated class from MongoDB: %v", err)
	}

	// Assert that the updated class fields match the expected updates.
	if updatedClass.Name != updatedClassInput.Name {
		t.Errorf("Updated class name does not match the expected value")
	}
	// Add similar assertions for other fields you updated.
}

func TestDeleteClassByID(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new ClassRepository using the test collection.
	repo := repository.NewClassRepository(collection)

	// Define your test data for an existing class.
	existingClass := &model.Class{
		ID:   "123",
		Name: "Test Class",
		// Initialize other fields as needed for the existing class.
	}

	// Insert the existing class into MongoDB.
	_, err = repo.CreateClass(existingClass)
	if err != nil {
		t.Errorf("Error creating the existing class: %v", err)
	}

	// Call the method you want to test.
	err = repo.DeleteClassByID("123")

	// Assert the error as needed.
	if err != nil {
		t.Errorf("Error deleting class: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()
	// Attempt to fetch the deleted class from MongoDB.
	filter := bson.M{"id": "123"}
	var databaseClass model.Class

	err = collection.FindOne(ctx, filter).Decode(&databaseClass)

	// Assert that the error is not nil, indicating the class was deleted.
	if err == nil {
		t.Errorf("Class was not deleted as expected")
	}
}

func TestGetClassByID(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new ClassRepository using the test collection.
	repo := repository.NewClassRepository(collection)

	// Define your test data for an existing class.
	existingClass := &model.Class{
		ID:   "123",
		Name: "Test Class",
		// Initialize other fields as needed for the existing class.
	}

	// Insert the existing class into MongoDB.
	_, err = repo.CreateClass(existingClass)
	if err != nil {
		t.Errorf("Error creating the existing class: %v", err)
	}

	// Call the method you want to test.
	fetchedClass, err := repo.GetClassByID("123")

	// Assert the result and error as needed.
	if err != nil {
		t.Errorf("Error fetching class: %v", err)
	}

	// Assert that the fetched class is not nil and has the expected ID.
	if fetchedClass == nil {
		t.Errorf("Fetched class is nil, expected a valid class")
	} else if fetchedClass.ID != "123" {
		t.Errorf("Fetched class has the wrong ID")
	}
}

func TestListClasses(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new ClassRepository using the test collection.
	repo := repository.NewClassRepository(collection)

	// Define your test data for multiple classes.
	class1 := &model.Class{
		ID:   "1",
		Name: "Class 1",
		// Initialize other fields as needed for class1.
	}
	class2 := &model.Class{
		ID:   "2",
		Name: "Class 2",
		// Initialize other fields as needed for class2.
	}

	// Insert the test classes into MongoDB.
	_, err = repo.CreateClass(class1)
	if err != nil {
		t.Errorf("Error creating class 1: %v", err)
	}
	_, err = repo.CreateClass(class2)
	if err != nil {
		t.Errorf("Error creating class 2: %v", err)
	}

	// Call the method you want to test.
	classes, err := repo.ListClasses()

	// Assert the result and error as needed.
	if err != nil {
		t.Errorf("Error listing classes: %v", err)
	}

	// Assert that the classes slice is not nil and contains the expected classes.
	if classes == nil {
		t.Errorf("List of classes is nil, expected a valid slice")
	} else if len(classes) != 2 {
		t.Errorf("Expected 2 classes, got %d", len(classes))
	}

	// Add specific assertions for each class in the list.
	if classes[0].ID != "1" || classes[1].ID != "2" {
		t.Errorf("Listed classes have incorrect IDs")
	}
}

func clearCollection(collection *mongo.Collection) error {
	_, err := collection.DeleteMany(context.TODO(), bson.D{})
	return err
}
