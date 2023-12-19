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

var adminToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJnWWlqam1Zd3Z5a2t3WUNlZUtpVzV3amxVM215dmVoNTRZSHlVZFc5MUFzIn0.eyJleHAiOjE3MDExMDgyMzEsImlhdCI6MTcwMTEwNzkzMSwianRpIjoiNzYxYTQwZjktNTMzMS00Mzc4LWI5OTktZjhjNWM3MGRkYWEzIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4ODg4L3JlYWxtcy9jbG91ZC1wcm9qZWN0IiwiYXVkIjpbInVzZXItbWFuYWdlbWVudC1jbGllbnQiLCJhY2NvdW50Il0sInN1YiI6IjQwN2VjMjNkLWM2ZjQtNDhkYi05YjFlLWZhN2Q3MDBmMjg2NiIsInR5cCI6IkJlYXJlciIsImF6cCI6ImxvZ2luLWNsaWVudCIsInNlc3Npb25fc3RhdGUiOiIwMWZhNjNkZi0wNDJmLTRmNTMtYmYzZi03NDNkYjFjMmY0MjYiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtY2xvdWQtcHJvamVjdCIsIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbInVwZGF0ZV9zY2hvb2wiLCJnZXRfY2xhc3NlcyIsImdldF9leGVyY2lzZXMiLCJkZWxldGVfbW9kdWxlIiwiZ2V0X3NjaG9vbHMiLCJkZWxldGVfZXhlcmNpc2UiLCJ1cGRhdGVfZXhlcmNpc2UiLCJnZXRfZXhlcmNpc2UiLCJkZWxldGVfbW9kdWxlX2FsbCIsImNyZWF0ZV9leGVyY2lzZSIsImdldF9zY2hvb2wiLCJkZWxldGVfZXhlcmNpc2VfYWxsIiwidXBkYXRlX3NjaG9vbF9hbGwiLCJkZWxldGVfY2xhc3MiLCJjcmVhdGVfbW9kdWxlIiwiZ2V0X21vZHVsZSIsInVwZGF0ZV9leGVyY2lzZV9hbGwiLCJnZXRfbW9kdWxlcyIsImNyZWF0ZV9jbGFzcyIsImNyZWF0ZV9zY2hvb2wiLCJ1cGRhdGVfbW9kdWxlX2FsbCIsImRlbGV0ZV9zY2hvb2wiLCJ1cGRhdGVfY2xhc3NfYWxsIiwidXBkYXRlX21vZHVsZSIsImdldF9jbGFzcyIsImRlbGV0ZV9zY2hvb2xfYWxsIiwidXBkYXRlX2NsYXNzIiwiZGVsZXRlX2NsYXNzX2FsbCJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwic2lkIjoiMDFmYTYzZGYtMDQyZi00ZjUzLWJmM2YtNzQzZGIxYzJmNDI2IiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiY2hhZCBhZG1pbiIsInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluQGFkbWluLmNvbSIsImdpdmVuX25hbWUiOiJjaGFkIiwiZmFtaWx5X25hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIn0.F4WBE9C3Ct17v5broRGPO92YR-lt9CzLprCnrOe4jWIMcMYyjHSBoLC-oQ7GHSoe1MjXe02CWRP98IZqQ5TPhF7nCliYs5qhn2vZRtlLa-QsjrTF2kZ1F_uEdXVekhVSKIRRFwoH8y2KxkaR3SSQ4bXOtJe8UJQs1AvzHPPeVDmRgfQcCZDNwdQTGI9Sb-8-C_dLXmU6W2ORJN1GmKikn9in4IS2kZ6KEiW6qNqOOllNlSQMZdtLQXf8BlymGf6s8z9j1itpg4iVljKeV8X76A8EHy-xQ98ESB188OVOxFHYReT82xOp5pusRjvMf3K71t20jPcOtUj-GiTZpEy9pQ"

func TestService_CreateSchool(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("CreateSchool", mock.AnythingOfType("string")).Return("", nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On("CreateSchool", mock.AnythingOfType("*model.School")).
		Return(&mocks.MockSchool, nil)

	result, err := service.CreateSchool(adminToken, mocks.MockCreateInput)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_CreateSchool_CatchValidationErrors(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("CreateSchool", mock.AnythingOfType("string")).Return("", nil)

	mockValidator.On("GetErrors").Return([]string{"validation_error"})

	result, err := service.CreateSchool(adminToken, mocks.MockCreateInput)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Validation errors: validation_error", err.Error())

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

func TestService_CreateSchool_CatchInsertError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("CreateSchool", mock.AnythingOfType("string")).Return("", nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On("CreateSchool", mock.AnythingOfType("*model.School")).
		Return(&model.School{}, errors.New("insertion_error"))

	result, err := service.CreateSchool(adminToken, mocks.MockCreateInput)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "insertion_error", err.Error())

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

func TestService_UpdateSchool(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("UpdateSchool", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.MockSchool, nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On(
			"UpdateSchool",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
			mock.AnythingOfType("model.School")).
		Return(&mocks.MockUpdatedSchool, nil)

	result, err := service.UpdateSchool(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockUpdateInput)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_UpdateSchool_CatchValidationErrors(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("UpdateSchool", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.MockSchool, nil)

	mockValidator.On("GetErrors").Return([]string{"validation_error"})

	result, err := service.UpdateSchool(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockUpdateInput)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Validation errors: validation_error", err.Error())

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

func TestService_UpdateSchool_CatchUpdateError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("UpdateSchool", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.MockSchool, nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On("UpdateSchool",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
			mock.AnythingOfType("model.School")).
		Return(&model.School{}, errors.New("update_error"))

	result, err := service.UpdateSchool(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockUpdateInput)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "update_error", err.Error())

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

func TestService_DeleteSchoolWithoutAdminToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("DeleteSchool", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.MockSchool, nil)

	mockRepo.
		On(
			"DeleteSchool",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mock.AnythingOfType("model.School")).
		Return(nil)

	err := service.DeleteSchool(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)

	mockPolicy.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_DeleteSchool_CatchDeleteError_WithoutAdminToken_AlreadySoftDeleted(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("DeleteSchool", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.SoftDeletedMockSchool, nil)

	err := service.DeleteSchool(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.NotNil(t, err)
	assert.Equal(t, "school could not be deleted", err.Error())

	mockPolicy.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_DeleteSchool_CatchDeleteError_WithAdminToken_AlreadySoftDeleted_NoFilter(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("DeleteSchool", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.SoftDeletedMockSchool, nil)

	err := service.DeleteSchool(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.NotNil(t, err)
	assert.Equal(t, "school could not be deleted", err.Error())

	mockPolicy.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_GetSchoolByID(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("GetSchool", mock.AnythingOfType("string")).Return(&mocks.MockSchool, nil)

	result, err := service.GetSchoolById(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)
}

func TestService_ListSchools(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("ListSchools", mock.AnythingOfType("string")).Return(nil)

	mockPolicy.On("HasPermissions", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(true)

	mockRepo.On("ListSchools").Return([]*model.SchoolInfo{&mocks.MockSchoolInfo}, nil)

	mockValidator.On("GetErrors").Return([]string{})

	filter := model.ListSchoolFilter{}
	paginate := model.Paginator{}
	result, err := service.ListSchools(adminToken, &filter, &paginate)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.IsType(t, &model.SchoolInfo{}, result[0])
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result[0].ID)

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_ListSchools_CatchRetrieveError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.SchoolService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("ListSchools", mock.AnythingOfType("string")).Return(nil)

	mockPolicy.On("HasPermissions", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(true)

	mockRepo.On("ListSchools").Return([]*model.SchoolInfo{}, errors.New("retrieval_error"))

	mockValidator.On("GetErrors").Return([]string{})

	filter := model.ListSchoolFilter{}
	paginate := model.Paginator{}
	result, err := service.ListSchools(adminToken, &filter, &paginate)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "retrieval_error", err.Error())

	mockPolicy.AssertExpectations(t)
}
