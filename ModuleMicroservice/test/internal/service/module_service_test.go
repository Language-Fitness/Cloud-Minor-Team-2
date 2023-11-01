package service

import (
	"Module/graph/model"
	service2 "Module/internal/service"
	"Module/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_CreateModule(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	newModule := model.ModuleInput{
		Name: "Test Module",
	}
	result, err := service.CreateModule(newModule)

	// Assert that there's no error
	assert.Nil(t, err)

	// Add assertions to verify that the result matches your expectations
	assert.NotNil(t, result)
	assert.Equal(t, "test-module-id", result.ID)

	// Assert that the expected functions were called on the mocks
	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}
