package service

import (
	service2 "example/internal/service"
	"example/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestService_CreateClass(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On("CreateClass", mock.AnythingOfType("*model.Module")).
		Return(&mocks.MockClass, nil)

	result, err := service.CreateClass(mocks.MockNewClass)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	//assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)
	//
	//mockValidator.AssertExpectations(t)
	//mockRepo.AssertExpectations(t)
}

//func TestService_CreateClass_CatchValidationErrors(t *testing.T) {
//	mockRepo := new(mocks.MockRepository)
//	mockValidator := new(mocks.MockValidator)
//	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo}
//
//	mockValidator.On("GetErrors").Return([]string{"validation_error"})
//
//	result, err := service.CreateClass(mocks.MockNewClass)
//
//	assert.NotNil(t, err)
//	assert.Nil(t, result)
//	assert.Equal(t, "Validation errors: validation_error", err.Error())
//
//	mockValidator.AssertExpectations(t)
//}
//
//func TestService_CreateClass_CatchInsertError(t *testing.T) {
//	mockRepo := new(mocks.MockRepository)
//	mockValidator := new(mocks.MockValidator)
//	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo}
//
//	mockValidator.On("GetErrors").Return([]string{})
//
//	mockRepo.
//		On("CreateClass", mock.AnythingOfType("*model.Module")).
//		Return(&model.Class{}, errors.New("insertion_error"))
//
//	result, err := service.CreateClass(mocks.MockNewClass)
//
//	assert.Nil(t, result)
//	assert.NotNil(t, err)
//	assert.Equal(t, "insertion_error", err.Error())
//
//	mockValidator.AssertExpectations(t)
//}
//
//func TestService_UpdateModule(t *testing.T) {
//	mockRepo := new(mocks.MockRepository)
//	mockValidator := new(mocks.MockValidator)
//	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo}
//
//	mockValidator.On("GetErrors").Return([]string{})
//
//	mockRepo.
//		On(
//			"UpdateModule",
//			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
//			mock.AnythingOfType("model.Module")).
//		Return(&mocks.MockUpdateClass, nil)
//
//	result, err := service.UpdateClass("3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockNewClass)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, result)
//	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)
//
//	mockValidator.AssertExpectations(t)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestService_UpdateModule_CatchValidationErrors(t *testing.T) {
//	mockRepo := new(mocks.MockRepository)
//	mockValidator := new(mocks.MockValidator)
//	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo}
//
//	mockValidator.On("GetErrors").Return([]string{"validation_error"})
//
//	result, err := service.UpdateClass("3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockNewClass)
//
//	assert.NotNil(t, err)
//	assert.Nil(t, result)
//	assert.Equal(t, "Validation errors: validation_error", err.Error())
//
//	mockValidator.AssertExpectations(t)
//}
//
//func TestService_UpdateModule_CatchUpdateError(t *testing.T) {
//	mockRepo := new(mocks.MockRepository)
//	mockValidator := new(mocks.MockValidator)
//	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo}
//
//	mockValidator.On("GetErrors").Return([]string{})
//
//	mockRepo.
//		On("UpdateModule",
//			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
//			mock.AnythingOfType("model.Module")).
//		Return(&model.Class{}, errors.New("update_error"))
//
//	result, err := service.UpdateClass("3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockNewClass)
//
//	assert.Nil(t, result)
//	assert.NotNil(t, err)
//	assert.Equal(t, "update_error", err.Error())
//
//	mockValidator.AssertExpectations(t)
//}
//
//func TestService_DeleteModule(t *testing.T) {
//	mockRepo := new(mocks.MockRepository)
//	mockValidator := new(mocks.MockValidator)
//	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo}
//
//	mockValidator.On("GetErrors").Return([]string{})
//
//	mockRepo.
//		On(
//			"DeleteModuleByID",
//			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
//		Return(nil)
//
//	err := service.DeleteClass("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")
//
//	assert.Nil(t, err)
//
//	mockValidator.AssertExpectations(t)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestService_DeleteModule_CatchValidationError(t *testing.T) {
//	mockRepo := new(mocks.MockRepository)
//	mockValidator := new(mocks.MockValidator)
//	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo}
//
//	mockValidator.On("GetErrors").Return([]string{"validation_error"})
//
//	err := service.DeleteClass("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")
//
//	assert.NotNil(t, err)
//	assert.Equal(t, "Validation errors: validation_error", err.Error())
//
//	mockValidator.AssertExpectations(t)
//}
//
//func TestService_DeleteModule_CatchDeleteError(t *testing.T) {
//	mockRepo := new(mocks.MockRepository)
//	mockValidator := new(mocks.MockValidator)
//	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo}
//
//	mockValidator.On("GetErrors").Return([]string{})
//
//	mockRepo.
//		On(
//			"DeleteModuleByID",
//			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
//		Return(errors.New("deletion_error"))
//
//	err := service.DeleteClass("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")
//
//	assert.NotNil(t, err)
//	assert.Equal(t, "deletion_error", err.Error())
//
//	mockValidator.AssertExpectations(t)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestService_GetModuleByID(t *testing.T) {
//	mockRepo := new(mocks.MockRepository)
//	mockValidator := new(mocks.MockValidator)
//	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo}
//
//	mockValidator.On("GetErrors").Return([]string{})
//
//	mockRepo.
//		On(
//			"GetModuleByID",
//			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
//		Return(&mocks.MockClass, nil)
//
//	result, err := service.GetClassById("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")
//
//	assert.Nil(t, err)
//	assert.NotNil(t, result)
//	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)
//
//	mockValidator.AssertExpectations(t)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestService_GetModuleByID_CatchValidationError(t *testing.T) {
//	mockRepo := new(mocks.MockRepository)
//	mockValidator := new(mocks.MockValidator)
//	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo}
//
//	mockValidator.On("GetErrors").Return([]string{"validation_error"})
//
//	result, err := service.GetClassById("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")
//
//	assert.Nil(t, result)
//	assert.NotNil(t, err)
//	assert.Equal(t, "Validation errors: validation_error", err.Error())
//
//	mockValidator.AssertExpectations(t)
//}
//
//func TestService_GetModuleByID_CatchRetrieveError(t *testing.T) {
//	mockRepo := new(mocks.MockRepository)
//	mockValidator := new(mocks.MockValidator)
//	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo}
//
//	mockValidator.On("GetErrors").Return([]string{})
//
//	mockRepo.
//		On(
//			"GetModuleByID",
//			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
//		Return(&model.Class{}, errors.New("retrieval_error"))
//
//	result, err := service.GetClassById("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")
//
//	assert.Nil(t, result)
//	assert.NotNil(t, err)
//	assert.Equal(t, "retrieval_error", err.Error())
//
//	mockValidator.AssertExpectations(t)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestService_ListModules(t *testing.T) {
//	mockRepo := new(mocks.MockRepository)
//	mockValidator := new(mocks.MockValidator)
//	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo}
//
//	mockRepo.On("ListModules").Return([]*model.Class{&mocks.MockClass}, nil)
//
//	result, err := service.ListClasses()
//
//	assert.Nil(t, err)
//	assert.NotNil(t, result)
//	assert.Len(t, result, 1)
//	assert.IsType(t, &model.Class{}, result[0])
//	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result[0].ID)
//
//	mockValidator.AssertExpectations(t)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestService_ListModules_CatchRetrieveError(t *testing.T) {
//	mockRepo := new(mocks.MockRepository)
//	mockValidator := new(mocks.MockValidator)
//	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo}
//
//	mockRepo.On("ListModules").Return([]*model.Class{}, errors.New("retrieval_error"))
//
//	result, err := service.ListClasses()
//
//	assert.Nil(t, result)
//	assert.NotNil(t, err)
//	assert.Equal(t, "retrieval_error", err.Error())
//}
