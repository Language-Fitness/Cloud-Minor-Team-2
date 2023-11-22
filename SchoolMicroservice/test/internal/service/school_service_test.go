package service

import (
	"errors"
	"example/graph/model"
	service2 "example/internal/service"
	"example/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestService_CreateSchool(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On("CreateSchool", mock.AnythingOfType("*model.School")).
		Return(&mocks.MockSchool, nil)

	result, err := service.CreateSchool(mocks.MockCreateInput)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_CreateSchool_CatchValidationErrors(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{"validation_error"})

	result, err := service.CreateSchool(mocks.MockCreateInput)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Validation errors: validation_error", err.Error())

	mockValidator.AssertExpectations(t)
}

func TestService_CreateSchool_CatchInsertError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On("CreateSchool", mock.AnythingOfType("*model.School")).
		Return(&model.School{}, errors.New("insertion_error"))

	result, err := service.CreateSchool(mocks.MockCreateInput)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "insertion_error", err.Error())

	mockValidator.AssertExpectations(t)
}

func TestService_UpdateSchool(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On(
			"GetSchoolByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockSchool, nil)

	mockRepo.
		On(
			"UpdateSchool",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
			mock.AnythingOfType("model.School")).
		Return(&mocks.MockUpdatedSchool, nil)

	result, err := service.UpdateSchool("3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockUpdateInput)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_UpdateSchool_CatchValidationErrors(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{"validation_error"})

	result, err := service.UpdateSchool("3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockUpdateInput)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Validation errors: validation_error", err.Error())

	mockValidator.AssertExpectations(t)
}

func TestService_UpdateSchool_CatchUpdateError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On(
			"GetSchoolByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockSchool, nil)

	mockRepo.
		On("UpdateSchool",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
			mock.AnythingOfType("model.School")).
		Return(&model.School{}, errors.New("update_error"))

	result, err := service.UpdateSchool("3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockUpdateInput)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "update_error", err.Error())

	mockValidator.AssertExpectations(t)
}

func TestService_DeleteSchool(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On(
			"DeleteSchoolByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(nil)

	err := service.DeleteSchool("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_DeleteSchool_CatchValidationError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{"validation_error"})

	err := service.DeleteSchool("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.NotNil(t, err)
	assert.Equal(t, "Validation errors: validation_error", err.Error())

	mockValidator.AssertExpectations(t)
}

func TestService_DeleteSchool_CatchDeleteError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On(
			"DeleteSchoolByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(errors.New("deletion_error"))

	err := service.DeleteSchool("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.NotNil(t, err)
	assert.Equal(t, "deletion_error", err.Error())

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_GetSchoolByID(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On(
			"GetSchoolByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockSchool, nil)

	result, err := service.GetSchoolById("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_GetSchoolByID_CatchValidationError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{"validation_error"})

	result, err := service.GetSchoolById("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "Validation errors: validation_error", err.Error())

	mockValidator.AssertExpectations(t)
}

func TestService_GetSchoolByID_CatchRetrieveError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo}

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On(
			"GetSchoolByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&model.School{}, errors.New("retrieval_error"))

	result, err := service.GetSchoolById("3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "retrieval_error", err.Error())

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_ListSchools(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo}

	mockRepo.On("ListSchools").Return([]*model.School{&mocks.MockSchool}, nil)

	result, err := service.ListSchools()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.IsType(t, &model.School{}, result[0])
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result[0].ID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_ListSchools_CatchRetrieveError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo}

	mockRepo.On("ListSchools").Return([]*model.School{}, errors.New("retrieval_error"))

	result, err := service.ListSchools()

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "retrieval_error", err.Error())
}
