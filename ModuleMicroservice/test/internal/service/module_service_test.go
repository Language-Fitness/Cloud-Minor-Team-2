package service

import (
	service2 "Module/internal/service"
	"Module/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestService_CreateModule(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})
	mockRepo.On("CreateModule", mock.AnythingOfType("*model.Module")).Return(&mocks.MockCreatedModule, nil)

	result, err := service.CreateModule(mocks.MockNewModule)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_CreateModule_GetErrors(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{"error1"})

	result, err := service.CreateModule(mocks.MockNewModule)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Validation errors: error1", err.Error())

	mockValidator.AssertExpectations(t)
}
