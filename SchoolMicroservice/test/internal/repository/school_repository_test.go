package graph

import (
	"context"
	"example/graph/model"
	"example/internal/repository"
	database "example/test/internal/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestCreateSchool(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new SchoolRepository using the test collection.
	repo := repository.NewSchoolRepository(collection)

	// Define your test data based on the school struct.
	location := "This is a sample location."
	timestamp := time.Now().String()
	softDeleted := false

	newSchool := &model.School{
		ID:          "123",
		Name:        "Test School",
		Location:    location,
		CreatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	}

	// Call the method you want to test.
	_, err = repo.CreateSchool(newSchool)

	// Assert the result and error as needed.
	if err != nil {
		t.Errorf("Error creating school: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()
	// Attempt to fetch the school from MongoDB.
	filter := bson.M{"id": "123"}
	var databaseSchool model.School

	err = collection.FindOne(ctx, filter).Decode(&databaseSchool)

	if err != nil {
		t.Errorf("Error fetching school from MongoDB: %v", err)
	}

	if !reflect.DeepEqual(newSchool, &databaseSchool) {
		t.Errorf("Retrieved school does not match the expected school")
	}
}

func TestUpdateSchool(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new SchoolRepository using the test collection.
	repo := repository.NewSchoolRepository(collection)

	// Define your test data for an existing school.
	existingSchool := &model.School{
		ID:   "123",
		Name: "Test School",
		// Initialize other fields as needed for the existing School.
	}

	// Insert the existing school into MongoDB.
	_, err = repo.CreateSchool(existingSchool)
	if err != nil {
		t.Errorf("Error creating the existing school: %v", err)
	}

	// Define the updates you want to apply to the school.
	updatedSchoolInput := model.School{
		ID:   "123",
		Name: "Updated Test School",
		// Define other fields you want to update.
	}

	// Call the method you want to test.
	updatedSchool, err := repo.UpdateSchool("123", updatedSchoolInput)

	// Assert the result and error as needed.
	if err != nil {
		t.Errorf("Error updating school: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()
	// Attempt to fetch the updated school from MongoDB.
	filter := bson.M{"id": "123"}
	var databaseSchool model.School

	err = collection.FindOne(ctx, filter).Decode(&databaseSchool)

	if err != nil {
		t.Errorf("Error fetching updated school from MongoDB: %v", err)
	}

	// Assert that the updated school fields match the expected updates.
	if updatedSchool.Name != updatedSchoolInput.Name {
		t.Errorf("Updated school name does not match the expected value")
	}
	// Add similar assertions for other fields you updated.
}

func TestSoftDeleteSchool(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new SchoolRepository using the test collection.
	repo := repository.NewSchoolRepository(collection)

	// Define your test data for an existing school.
	isNotSoftDelete := false
	existingSchool := &model.School{
		ID:          "123",
		Name:        "Test School",
		SoftDeleted: &isNotSoftDelete,
		// Initialize other fields as needed for the existing School.
	}

	// Insert the existing school into MongoDB.
	_, err = repo.CreateSchool(existingSchool)
	if err != nil {
		t.Errorf("Error creating the existing school: %v", err)
	}

	// Define the updates you want to apply to the school.
	updatedSchoolInput := model.School{
		ID:          "123",
		Name:        "Updated Test School",
		SoftDeleted: &isNotSoftDelete,
		// Define other fields you want to update.
	}

	// Call the method you want to test.
	updatedSchool, err := repo.UpdateSchool("123", updatedSchoolInput)

	// Assert the result and error as needed.
	if err != nil {
		t.Errorf("Error updating school: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()
	// Attempt to fetch the updated school from MongoDB.
	filter := bson.M{"id": "123"}
	var databaseSchool model.School

	err = collection.FindOne(ctx, filter).Decode(&databaseSchool)

	if err != nil {
		t.Errorf("Error fetching updated school from MongoDB: %v", err)
	}

	// Assert that the updated school fields match the expected updates.
	if *updatedSchool.SoftDeleted != *updatedSchoolInput.SoftDeleted {
		t.Errorf("Updated school name does not match the expected value")
	}
	// Add similar assertions for other fields you updated.
}

func TestDeleteSchoolByID(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new SchoolRepository using the test collection.
	repo := repository.NewSchoolRepository(collection)

	// Define your test data for an existing school.
	existingSchool := &model.School{
		ID:   "123",
		Name: "Test School",
		// Initialize other fields as needed for the existing school.
	}

	// Insert the existing school into MongoDB.
	_, err = repo.CreateSchool(existingSchool)
	if err != nil {
		t.Errorf("Error creating the existing school: %v", err)
	}

	// Call the method you want to test.
	err = repo.HardDeleteSchoolByID("123")

	// Assert the error as needed.
	if err != nil {
		t.Errorf("Error deleting school: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()
	// Attempt to fetch the deleted school from MongoDB.
	filter := bson.M{"id": "123"}
	var databaseSchool model.School

	err = collection.FindOne(ctx, filter).Decode(&databaseSchool)

	// Assert that the error is not nil, indicating the school was deleted.
	if err == nil {
		t.Errorf("school was not deleted as expected")
	}
}

func TestGetSchoolByID(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new SchoolRepository using the test collection.
	repo := repository.NewSchoolRepository(collection)

	// Define your test data for an existing school.
	existingSchool := &model.School{
		ID:   "123",
		Name: "Test School",
		// Initialize other fields as needed for the existing school.
	}

	// Insert the existing school into MongoDB.
	_, err = repo.CreateSchool(existingSchool)
	if err != nil {
		t.Errorf("Error creating the existing school: %v", err)
	}

	// Call the method you want to test.
	fetchedSchool, err := repo.GetSchoolByID("123")

	// Assert the result and error as needed.
	if err != nil {
		t.Errorf("Error fetching school: %v", err)
	}

	// Assert that the fetched school is not nil and has the expected ID.
	if fetchedSchool == nil {
		t.Errorf("Fetched school is nil, expected a valid school")
	} else if fetchedSchool.ID != "123" {
		t.Errorf("Fetched school has the wrong ID")
	}
}

func TestListSchools(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		log.Fatalf("Failed to clear the test collection: %v", err)
	}

	// Create a new SchoolRepository using the test collection.
	repo := repository.NewSchoolRepository(collection)

	// Define your test data for multiple schools.
	school1 := &model.School{
		ID:   "1",
		Name: "School 1",
		// Initialize other fields as needed for school1.
	}
	school2 := &model.School{
		ID:   "2",
		Name: "School 2",
		// Initialize other fields as needed for school2.
	}

	// Insert the test schools into MongoDB.
	_, err = repo.CreateSchool(school1)
	if err != nil {
		t.Errorf("Error creating school 1: %v", err)
	}
	_, err = repo.CreateSchool(school2)
	if err != nil {
		t.Errorf("Error creating school 2: %v", err)
	}

	paginateOptions := options.Find().
		SetSkip(int64(0)).
		SetLimit(int64(2))

	// Call the method you want to test.
	schools, err := repo.ListSchools(bson.D{}, paginateOptions)

	// Assert the result and error as needed.
	if err != nil {
		t.Errorf("Error listing schooles: %v", err)
	}

	// Assert that the schools slice is not nil and contains the expected schools.
	if schools == nil {
		t.Errorf("List of schooles is nil, expected a valid slice")
	} else if len(schools) != 2 {
		t.Errorf("Expected 2 schooles, got %d", len(schools))
	}

	// Add specific assertions for each school in the list.
	if schools[0].ID != "1" || schools[1].ID != "2" {
		t.Errorf("Listed schools have incorrect IDs")
	}
}

func clearCollection(collection *mongo.Collection) error {
	_, err := collection.DeleteMany(context.TODO(), bson.D{})
	return err
}
