package graph

import (
	"ExerciseMicroservice/graph"
	"ExerciseMicroservice/graph/model"
	"ExerciseMicroservice/test/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateExercise(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockExerciseService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service
	mockService.On("CreateExercise", mock.Anything, mocks.MockExerciseInput).Return(&mocks.MockExercise, nil).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	exercise, err := resolver.Mutation().CreateExercise(ctx, mocks.MockExerciseInput)

	// Assert that the result and error match expectations
	assert.NoError(t, err)
	assert.Equal(t, &mocks.MockExercise, exercise)

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestCreateExercise_Error(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockExerciseService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service for an error case
	expectedError := errors.New("some error")
	mockService.On("CreateExercise", mock.Anything, mocks.MockExerciseInput).Return((*model.Exercise)(nil), expectedError).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	exercise, err := resolver.Mutation().CreateExercise(ctx, mocks.MockExerciseInput)

	// Assert that the result is nil and the error is not nil
	assert.Nil(t, exercise)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestUpdateExercise(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockExerciseService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service
	mockService.On("UpdateExercise", mock.Anything, "exerciseID", mocks.MockExerciseInput).Return(&mocks.MockExercise, nil).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	exercise, err := resolver.Mutation().UpdateExercise(ctx, "exerciseID", mocks.MockExerciseInput)

	// Assert that the result and error match expectations
	assert.NoError(t, err)
	assert.Equal(t, &mocks.MockExercise, exercise)

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestUpdateExercise_Error(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockExerciseService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service for an error case
	expectedError := errors.New("some error")
	mockService.On("UpdateExercise", mock.Anything, "exerciseID", mocks.MockExerciseInput).Return((*model.Exercise)(nil), expectedError).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	exercise, err := resolver.Mutation().UpdateExercise(ctx, "exerciseID", mocks.MockExerciseInput)

	// Assert that the result is nil and the error is not nil
	assert.Nil(t, exercise)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestGetExercise(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockExerciseService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service
	mockService.On("GetExerciseById", mock.Anything, "exerciseID").Return(&mocks.MockExercise, nil).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	exercise, err := resolver.Query().GetExercise(ctx, "exerciseID")

	// Assert that the result and error match expectations
	assert.NoError(t, err)
	assert.Equal(t, &mocks.MockExercise, exercise)

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestGetExercise_Error(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockExerciseService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service for an error case
	expectedError := errors.New("some error")
	mockService.On("GetExerciseById", mock.Anything, "exerciseID").Return((*model.Exercise)(nil), expectedError).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	exercise, err := resolver.Query().GetExercise(ctx, "exerciseID")

	// Assert that the result is nil and the error is not nil
	assert.Nil(t, exercise)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestListExercise(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockExerciseService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service
	mockService.On("ListExercises", mock.Anything, mock.Anything, mock.Anything).Return([]*model.ExerciseInfo{&mocks.MockExerciseInfo}, nil).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	exercises, err := resolver.Query().ListExercise(ctx, model.ExerciseFilter{}, model.Paginator{})

	// Assert that the result and error match expectations
	assert.NoError(t, err)
	assert.Equal(t, []*model.ExerciseInfo{&mocks.MockExerciseInfo}, exercises)

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestListExercise_Error(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockExerciseService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service for an error case
	expectedError := errors.New("some error")
	mockService.On("ListExercises", mock.Anything, mock.Anything, mock.Anything).Return([]*model.ExerciseInfo{(*model.ExerciseInfo)(nil)}, expectedError).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	exercises, err := resolver.Query().ListExercise(ctx, model.ExerciseFilter{}, model.Paginator{})

	// Assert that the result is nil and the error is not nil
	assert.Nil(t, exercises)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}
