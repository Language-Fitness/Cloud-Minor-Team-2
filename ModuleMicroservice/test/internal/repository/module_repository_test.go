package graph

import (
	"Module/graph/model"
	"Module/internal/repository"
	database "Module/test/internal/helpers"
	"testing"
)

func TestCreateModule(t *testing.T) {
	// Get the MongoDB collection for testing.
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
	result, err := repo.CreateModule(newModule)

	// Assert the result and error as needed.
	if err != nil {
		t.Errorf("Error creating module: %v", err)
	}

	println(result.ID)
}
