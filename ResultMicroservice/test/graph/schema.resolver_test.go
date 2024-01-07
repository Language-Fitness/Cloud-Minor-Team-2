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
	mockService.On("CreateResult", mock.Anything, mocks.MockResultInput).Return(&mocks.MockResult, nil).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	result, err := resolver.Mutation().CreateResult(ctx, mocks.MockResultInput)

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
	mockService.On("CreateResult", mock.Anything, mocks.MockResultInput).Return((*model.Result)(nil), expectedError).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	result, err := resolver.Mutation().CreateResult(ctx, mocks.MockResultInput)

	// Assert that the result is nil and the error is not nil
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestUpdateResult(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockResultService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service
	mockService.On("UpdateResult", mock.Anything, "resultID", mocks.MockResultInput).Return(&mocks.MockResult, nil).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	result, err := resolver.Mutation().UpdateResult(ctx, "resultID", mocks.MockResultInput)

	// Assert that the result and error match expectations
	assert.NoError(t, err)
	assert.Equal(t, &mocks.MockResult, result)

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestUpdateResult_Error(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockResultService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service for an error case
	expectedError := errors.New("some error")
	mockService.On("UpdateResult", mock.Anything, "resultID", mocks.MockResultInput).Return((*model.Result)(nil), expectedError).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	result, err := resolver.Mutation().UpdateResult(ctx, "resultID", mocks.MockResultInput)

	// Assert that the result is nil and the error is not nil
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestDeleteResult(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockResultService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service
	mockService.On("DeleteResult", mock.Anything, "resultID").Return(&mocks.MockResult, nil).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	result, err := resolver.Mutation().DeleteResult(ctx, "resultID")

	// Assert that the result and error match expectations
	assert.NoError(t, err)
	assert.Equal(t, &mocks.MockResult, result)

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestDeleteResult_Error(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockResultService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service for an error case
	expectedError := errors.New("some error")
	mockService.On("DeleteResult", mock.Anything, "resultID").Return((*model.Result)(nil), expectedError).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	result, err := resolver.Mutation().DeleteResult(ctx, "resultID")

	// Assert that the result is nil and the error is not nil
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestListResults(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockResultService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service
	mockService.On("ListResults", mock.Anything, mock.Anything, mock.Anything).Return([]*model.Result{&mocks.MockResult}, nil).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	results, err := resolver.Query().ListResults(ctx, model.ResultFilter{}, model.Paginator{})

	// Assert that the result and error match expectations
	assert.NoError(t, err)
	assert.Equal(t, []*model.Result{&mocks.MockResult}, results)

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestListResults_Error(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockResultService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service for an error case
	expectedError := errors.New("some error")
	mockService.On("ListResults", mock.Anything, mock.Anything, mock.Anything).Return([]*model.Result{(*model.Result)(nil)}, expectedError).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	results, err := resolver.Query().ListResults(ctx, model.ResultFilter{}, model.Paginator{})

	// Assert that the result is nil and the error is not nil
	assert.Nil(t, results)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestGetResultsByID(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockResultService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service
	mockService.On("GetResultById", mock.Anything, "resultID").Return(&mocks.MockResult, nil).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	result, err := resolver.Query().GetResultsByID(ctx, "resultID")

	// Assert that the result and error match expectations
	assert.NoError(t, err)
	assert.Equal(t, &mocks.MockResult, result)

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}

func TestGetResultsByID_Error(t *testing.T) {
	// Initialize mock service
	mockService := &mocks.MockResultService{}

	// Create a resolver with the mock service
	resolver := graph.Resolver{Service: mockService}

	// Set up expectations on the mock service for an error case
	expectedError := errors.New("some error")
	mockService.On("GetResultById", mock.Anything, "resultID").Return((*model.Result)(nil), expectedError).Once()

	// Create a context with a mock token
	ctx := context.WithValue(context.Background(), "myCustomTokenKey", "mock_token")

	// Call the resolver function
	result, err := resolver.Query().GetResultsByID(ctx, "resultID")

	// Assert that the result is nil and the error is not nil
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())

	// Verify that the expectations on the mock service were met
	mockService.AssertExpectations(t)
}
