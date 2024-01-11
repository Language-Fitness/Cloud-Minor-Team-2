package service

import (
	"ExerciseMicroservice/graph/model"
	"ExerciseMicroservice/internal/service"
	"ExerciseMicroservice/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var userToken = "user_token"

func TestExerciseService_CreateExercise(t *testing.T) {
	// Setup mock objects
	mockValidator := new(mocks.MockValidator)
	mockRepo := new(mocks.MockExerciseRepository)
	mockPolicy := new(mocks.MockExercisePolicy)

	// Create an instance of ExerciseService with mock dependencies
	exerciseService := service.ExerciseService{
		Validator: mockValidator,
		Repo:      mockRepo,
		Policy:    mockPolicy,
	}

	// Mock Policy expectations
	mockPolicy.On("CreateExercise", mock.AnythingOfType("string")).Return("", nil)

	// Mock Validator expectations
	mockValidator.On("GetErrors").Return([]string{})

	// Mock Repo expectations
	mockRepo.On("CreateExercise", mock.AnythingOfType("*model.Exercise")).Return(&model.Exercise{ID: mocks.ExerciseID}, nil)

	// Call the method being tested
	result, err := exerciseService.CreateExercise(userToken, mocks.MockExerciseInput)

	// Assertions
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mocks.ExerciseID, result.ID)

	// Verify that the expected methods were called
	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestExerciseService_UpdateExercise(t *testing.T) {
	// Setup mock objects
	mockValidator := new(mocks.MockValidator)
	mockRepo := new(mocks.MockExerciseRepository)
	mockPolicy := new(mocks.MockExercisePolicy)

	// Create an instance of ExerciseService with mock dependencies
	exerciseService := service.ExerciseService{
		Validator: mockValidator,
		Repo:      mockRepo,
		Policy:    mockPolicy,
	}

	// Mock Policy expectations
	mockPolicy.On("UpdateExercise", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&model.Exercise{}, nil)

	// Mock Validator expectations
	mockValidator.On("GetErrors").Return([]string{})

	// Mock Repo expectations
	mockRepo.On("UpdateExercise", mock.AnythingOfType("string"), mock.AnythingOfType("model.Exercise")).
		Return(&model.Exercise{ID: mocks.ExerciseID}, nil)

	// Call the method being tested
	result, err := exerciseService.UpdateExercise(userToken, mocks.ExerciseID, mocks.MockExerciseInput)

	// Assertions
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mocks.ExerciseID, result.ID)

	// Verify that the expected methods were called
	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestExerciseService_DeleteExercise(t *testing.T) {
	// Setup mock objects
	mockValidator := new(mocks.MockValidator)
	mockRepo := new(mocks.MockExerciseRepository)
	mockPolicy := new(mocks.MockExercisePolicy)

	// Create an instance of ExerciseService with mock dependencies
	exerciseService := service.ExerciseService{
		Validator: mockValidator,
		Repo:      mockRepo,
		Policy:    mockPolicy,
	}

	// Mock Validator expectations
	mockValidator.On("GetErrors").Return([]string{})

	// Mock Policy expectations
	mockPolicy.On("DeleteExercise", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(true, &model.Exercise{}, nil)

	// Mock Repo expectations
	mockRepo.On("UpdateExercise", mock.AnythingOfType("string"), mock.AnythingOfType("model.Exercise")).
		Return(&mocks.MockDeletedExercise, nil)

	// Call the method being tested
	err := exerciseService.DeleteExercise(userToken, mocks.ExerciseID, true)

	// Assertions
	assert.Nil(t, err)

	// Verify that the expected methods were called
	mockPolicy.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestExerciseService_UnDeleteExercise(t *testing.T) {
	// Setup mock objects
	mockValidator := new(mocks.MockValidator)
	mockRepo := new(mocks.MockExerciseRepository)
	mockPolicy := new(mocks.MockExercisePolicy)

	// Create an instance of ExerciseService with mock dependencies
	exerciseService := service.ExerciseService{
		Validator: mockValidator,
		Repo:      mockRepo,
		Policy:    mockPolicy,
	}

	// Mock Validator expectations
	mockValidator.On("GetErrors").Return([]string{})

	// Mock Policy expectations
	mockPolicy.On("DeleteExercise", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(true, &mocks.MockDeletedExercise, nil)

	// Mock Repo expectations
	mockRepo.On("UpdateExercise", mock.AnythingOfType("string"), mock.AnythingOfType("model.Exercise")).
		Return(&mocks.MockExercise, nil)

	// Call the method being tested
	err := exerciseService.DeleteExercise(userToken, mocks.ExerciseID, false)

	// Assertions
	assert.Nil(t, err)

	// Verify that the expected methods were called
	mockPolicy.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestExerciseService_GetExerciseById(t *testing.T) {
	// Setup mock objects
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockExercisePolicy)
	mockRepo := new(mocks.MockExerciseRepository)

	// Create an instance of ExerciseService with mock dependencies
	exerciseService := service.ExerciseService{
		Validator: mockValidator,
		Repo:      mockRepo,
		Policy:    mockPolicy,
	}

	// Mock Validator expectations
	mockValidator.On("GetErrors").Return([]string{})

	// Mock Policy expectations
	mockPolicy.On("GetExercise", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&model.Exercise{ID: mocks.ExerciseID}, nil)

	// Call the method being tested
	result, err := exerciseService.GetExerciseById(userToken, mocks.ExerciseID)

	// Assertions
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mocks.ExerciseID, result.ID)

	// Verify that the expected methods were called
	mockPolicy.AssertExpectations(t)
}

func TestExerciseService_ListExercises(t *testing.T) {
	// Setup mock objects
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockExercisePolicy)
	mockRepo := new(mocks.MockExerciseRepository)

	// Create an instance of ExerciseService with mock dependencies
	exerciseService := service.ExerciseService{
		Validator: mockValidator,
		Repo:      mockRepo,
		Policy:    mockPolicy,
	}

	// Mock Policy expectations
	mockPolicy.On("ListExercises", mock.AnythingOfType("string")).Return(true, nil)

	// Mock Validator expectations
	mockValidator.On("GetErrors").Return([]string{})

	// Mock Repo expectations
	mockRepo.On("ListExercises", mock.AnythingOfType("primitive.D"), mock.AnythingOfType("*options.FindOptions")).
		Return([]*model.ExerciseInfo{&mocks.MockExerciseInfo}, nil)

	// Call the method being tested
	result, err := exerciseService.ListExercises(userToken, &model.ExerciseFilter{}, &model.Paginator{})

	// Assertions
	assert.Nil(t, err)
	assert.NotNil(t, result)

	// Verify that the expected methods were called
	mockPolicy.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockValidator.AssertExpectations(t)

}
