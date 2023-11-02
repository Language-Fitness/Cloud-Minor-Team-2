package service

import (
	"Module/graph/model"
	service2 "Module/internal/service"
	"Module/test/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestService_CreateModule(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On("CreateModule", mock.AnythingOfType("*model.Module")).
		Return(&mocks.MockModule, nil)

	result, err := service.CreateModule(mocks.MockNewModule)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_CreateModule_CatchValidationErrors(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{"validation_error"})

	result, err := service.CreateModule(mocks.MockNewModule)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Validation errors: validation_error", err.Error())

	mockValidator.AssertExpectations(t)
}

func TestService_CreateModule_CatchInsertError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On("CreateModule", mock.AnythingOfType("*model.Module")).
		Return(&model.Module{}, errors.New("insertion_error"))

	result, err := service.CreateModule(mocks.MockNewModule)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "insertion_error", err.Error())

	mockValidator.AssertExpectations(t)
}

func TestService_UpdateModule(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On(
			"UpdateModule",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
			mock.AnythingOfType("model.Module")).
		Return(&mocks.MockUpdateModule, nil)

	result, err := service.UpdateModule("3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockUpdateModule)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_UpdateModule_CatchValidationErrors(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{"validation_error"})

	result, err := service.UpdateModule("3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockUpdateModule)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Validation errors: validation_error", err.Error())

	mockValidator.AssertExpectations(t)
}

func TestService_UpdateModule_CatchUpdateError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On("UpdateModule",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
			mock.AnythingOfType("model.Module")).
		Return(&model.Module{}, errors.New("update_error"))

	result, err := service.UpdateModule("3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockUpdateModule)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "update_error", err.Error())

	mockValidator.AssertExpectations(t)
}

func TestService_DeleteModule(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On(
			"DeleteModuleByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(nil)

	err := service.DeleteModule("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_DeleteModule_CatchValidationError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{"validation_error"})

	err := service.DeleteModule("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.NotNil(t, err)
	assert.Equal(t, "Validation errors: validation_error", err.Error())

	mockValidator.AssertExpectations(t)
}

func TestService_DeleteModule_CatchDeleteError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On(
			"DeleteModuleByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(errors.New("deletion_error"))

	err := service.DeleteModule("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.NotNil(t, err)
	assert.Equal(t, "deletion_error", err.Error())

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_GetModuleByID(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On(
			"GetModuleByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockModule, nil)

	result, err := service.GetModuleById("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_GetModuleByID_CatchValidationError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{"validation_error"})

	result, err := service.GetModuleById("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "Validation errors: validation_error", err.Error())

	mockValidator.AssertExpectations(t)
}

func TestService_GetModuleByID_CatchRetrieveError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On(
			"GetModuleByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&model.Module{}, errors.New("retrieval_error"))

	result, err := service.GetModuleById("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "retrieval_error", err.Error())

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_ListModules(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockRepo.On("ListModules").Return([]*model.Module{&mocks.MockModule}, nil)

	result, err := service.ListModules()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.IsType(t, &model.Module{}, result[0])
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result[0].ID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_ListModules_CatchRetrieveError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ModuleService{Validator: mockValidator, Repo: mockRepo}

	mockRepo.On("ListModules").Return([]*model.Module{}, errors.New("retrieval_error"))

	result, err := service.ListModules()

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "retrieval_error", err.Error())
}
