package repository

import (
	"ExerciseMicroservice/graph/model"
	"ExerciseMicroservice/internal/repository"
	database "ExerciseMicroservice/test/internal/helpers"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"testing"
	"time"
)

func TestCreateExercise(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		t.Fatalf("Failed to clear the test collection: %v", err)
	}

	repo := repository.NewExerciseRepository(collection)

	newExercise := &model.Exercise{
		ID:         "123",
		Name:       "Test Exercise",
		Difficulty: model.LanguageLevelA2,
		CreatedAt:  time.Now().String(),
	}

	createdExercise, err := repo.CreateExercise(newExercise)

	if err != nil {
		t.Errorf("Error creating exercise: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"id": "123"}
	var databaseExercise model.Exercise

	err = collection.FindOne(ctx, filter).Decode(&databaseExercise)

	if err != nil {
		t.Errorf("Error fetching exercise from MongoDB: %v", err)
	}

	if !reflect.DeepEqual(createdExercise, &databaseExercise) {
		t.Errorf("Retrieved exercise does not match the expected exercise")
	}
	fmt.Println("TestCreateExercise passed")
}

func TestUpdateExercise(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		t.Fatalf("Failed to clear the test collection: %v", err)
	}

	repo := repository.NewExerciseRepository(collection)

	existingExercise := &model.Exercise{
		ID:         "123",
		Name:       "Test Exercise",
		Difficulty: model.LanguageLevelB1,
		CreatedAt:  time.Now().String(),
	}

	_, err = repo.CreateExercise(existingExercise)
	if err != nil {
		t.Errorf("Error creating the existing exercise: %v", err)
	}

	updatedExerciseInput := model.Exercise{
		ID:         "123",
		Name:       "Updated Test Exercise",
		Difficulty: model.LanguageLevelA2,
		CreatedAt:  time.Now().String(),
	}

	updatedExercise, err := repo.UpdateExercise("123", updatedExerciseInput)

	if err != nil {
		t.Errorf("Error updating exercise: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"id": "123"}
	var databaseExercise model.Exercise

	err = collection.FindOne(ctx, filter).Decode(&databaseExercise)

	if err != nil {
		t.Errorf("Error fetching updated exercise from MongoDB: %v", err)
	}

	if updatedExercise.Name != updatedExerciseInput.Name {
		t.Errorf("Updated exercise name does not match the expected value")
	}
	fmt.Println("TestUpdateExercise passed")
}

func TestGetExerciseByID(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		t.Fatalf("Failed to clear the test collection: %v", err)
	}

	repo := repository.NewExerciseRepository(collection)

	existingExercise := &model.Exercise{
		ID:         "123",
		Name:       "Test Exercise",
		Difficulty: model.LanguageLevelA2,
		CreatedAt:  time.Now().String(),
	}

	_, err = repo.CreateExercise(existingExercise)
	if err != nil {
		t.Errorf("Error creating the existing exercise: %v", err)
	}

	fetchedExercise, err := repo.GetExerciseByID("123")

	if err != nil {
		t.Errorf("Error fetching exercise: %v", err)
	}

	if fetchedExercise == nil {
		t.Errorf("Fetched exercise is nil, expected a valid exercise")
	} else if fetchedExercise.ID != "123" {
		t.Errorf("Fetched exercise has the wrong ID")
	}
	fmt.Println("TestGetExerciseByID passed")
}

func TestListExercises(t *testing.T) {
	collection, err := database.GetTestCollection()
	if err != nil {
		t.Fatalf("Failed to get the test collection: %v", err)
	}

	err = clearCollection(collection)
	if err != nil {
		t.Fatalf("Failed to clear the test collection: %v", err)
	}

	repo := repository.NewExerciseRepository(collection)

	exercise1 := &model.Exercise{
		ID:         "1",
		Name:       "Exercise 1",
		Difficulty: model.LanguageLevelA2,
		CreatedAt:  time.Now().String(),
	}
	exercise2 := &model.Exercise{
		ID:         "2",
		Name:       "Exercise 2",
		Difficulty: model.LanguageLevelB1,
		CreatedAt:  time.Now().String(),
	}

	_, err = repo.CreateExercise(exercise1)
	if err != nil {
		t.Errorf("Error creating exercise 1: %v", err)
	}
	_, err = repo.CreateExercise(exercise2)
	if err != nil {
		t.Errorf("Error creating exercise 2: %v", err)
	}

	paginateOptions := options.Find().
		SetSkip(int64(0)).
		SetLimit(int64(2))

	exercises, err := repo.ListExercises(bson.D{}, paginateOptions)

	if err != nil {
		t.Errorf("Error listing exercises: %v", err)
	}

	if exercises == nil {
		t.Errorf("List of exercises is nil, expected a valid slice")
	} else if len(exercises) != 2 {
		t.Errorf("Expected 2 exercises, got %d", len(exercises))
	}

	if exercises[0].ID != "1" || exercises[1].ID != "2" {
		t.Errorf("Listed exercises have incorrect IDs")
	}
	fmt.Println("TestListExercises passed")
}

func clearCollection(collection *mongo.Collection) error {
	_, err := collection.DeleteMany(context.TODO(), bson.D{})
	return err
}
