package graph

import (
	"ResultMicroservice/graph"
	"ResultMicroservice/graph/model"
	"ResultMicroservice/test/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateResult(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockResultService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service
	mockService.On("CreateResult", mock.Anything, mocks.MockInputResult).Return(&mocks.MockResult, nil).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	result, err := resolver.Mutation().CreateResult(ctx, mocks.MockInputResult)

	// Assert that the result and error match expectations
	assert.NoError(t, err)
	assert.Equal(t, &mocks.MockResult, result)

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestCreateResult_Error(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockResultService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service for an error case
	expectedError := errors.New("some error")
	mockService.On("CreateResult", mock.Anything, mocks.MockInputResult).Return(&model.Result{}, expectedError).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	result, err := resolver.Mutation().CreateResult(ctx, mocks.MockInputResult)

	// Assert that the result is nil and the error is not nil
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}
