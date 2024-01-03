package service

import (
	"Class/graph/model"
	service2 "Class/internal/service"
	"Class/test/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var adminToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJnWWlqam1Zd3Z5a2t3WUNlZUtpVzV3amxVM215dmVoNTRZSHlVZFc5MUFzIn0.eyJleHAiOjE3MDExMDgyMzEsImlhdCI6MTcwMTEwNzkzMSwianRpIjoiNzYxYTQwZjktNTMzMS00Mzc4LWI5OTktZjhjNWM3MGRkYWEzIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4ODg4L3JlYWxtcy9jbG91ZC1wcm9qZWN0IiwiYXVkIjpbInVzZXItbWFuYWdlbWVudC1jbGllbnQiLCJhY2NvdW50Il0sInN1YiI6IjQwN2VjMjNkLWM2ZjQtNDhkYi05YjFlLWZhN2Q3MDBmMjg2NiIsInR5cCI6IkJlYXJlciIsImF6cCI6ImxvZ2luLWNsaWVudCIsInNlc3Npb25fc3RhdGUiOiIwMWZhNjNkZi0wNDJmLTRmNTMtYmYzZi03NDNkYjFjMmY0MjYiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtY2xvdWQtcHJvamVjdCIsIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbInVwZGF0ZV9zY2hvb2wiLCJnZXRfY2xhc3NlcyIsImdldF9leGVyY2lzZXMiLCJkZWxldGVfbW9kdWxlIiwiZ2V0X3NjaG9vbHMiLCJkZWxldGVfZXhlcmNpc2UiLCJ1cGRhdGVfZXhlcmNpc2UiLCJnZXRfZXhlcmNpc2UiLCJkZWxldGVfbW9kdWxlX2FsbCIsImNyZWF0ZV9leGVyY2lzZSIsImdldF9zY2hvb2wiLCJkZWxldGVfZXhlcmNpc2VfYWxsIiwidXBkYXRlX3NjaG9vbF9hbGwiLCJkZWxldGVfY2xhc3MiLCJjcmVhdGVfbW9kdWxlIiwiZ2V0X21vZHVsZSIsInVwZGF0ZV9leGVyY2lzZV9hbGwiLCJnZXRfbW9kdWxlcyIsImNyZWF0ZV9jbGFzcyIsImNyZWF0ZV9zY2hvb2wiLCJ1cGRhdGVfbW9kdWxlX2FsbCIsImRlbGV0ZV9zY2hvb2wiLCJ1cGRhdGVfY2xhc3NfYWxsIiwidXBkYXRlX21vZHVsZSIsImdldF9jbGFzcyIsImRlbGV0ZV9zY2hvb2xfYWxsIiwidXBkYXRlX2NsYXNzIiwiZGVsZXRlX2NsYXNzX2FsbCJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwic2lkIjoiMDFmYTYzZGYtMDQyZi00ZjUzLWJmM2YtNzQzZGIxYzJmNDI2IiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiY2hhZCBhZG1pbiIsInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluQGFkbWluLmNvbSIsImdpdmVuX25hbWUiOiJjaGFkIiwiZmFtaWx5X25hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIn0.F4WBE9C3Ct17v5broRGPO92YR-lt9CzLprCnrOe4jWIMcMYyjHSBoLC-oQ7GHSoe1MjXe02CWRP98IZqQ5TPhF7nCliYs5qhn2vZRtlLa-QsjrTF2kZ1F_uEdXVekhVSKIRRFwoH8y2KxkaR3SSQ4bXOtJe8UJQs1AvzHPPeVDmRgfQcCZDNwdQTGI9Sb-8-C_dLXmU6W2ORJN1GmKikn9in4IS2kZ6KEiW6qNqOOllNlSQMZdtLQXf8BlymGf6s8z9j1itpg4iVljKeV8X76A8EHy-xQ98ESB188OVOxFHYReT82xOp5pusRjvMf3K71t20jPcOtUj-GiTZpEy9pQ"

func TestService_CreateClass(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("CreateClass", mock.AnythingOfType("string")).Return("", nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On("CreateClass", mock.AnythingOfType("*model.Class")).
		Return(&mocks.MockClass, nil)

	result, err := service.CreateClass(adminToken, mocks.MockCreateInput)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_CreateClass_CatchValidationErrors(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("CreateClass", mock.AnythingOfType("string")).Return("", nil)

	mockValidator.On("GetErrors").Return([]string{"validation_error"})

	result, err := service.CreateClass(adminToken, mocks.MockCreateInput)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Validation errors: validation_error", err.Error())

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

func TestService_CreateClass_CatchInsertError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("CreateClass", mock.AnythingOfType("string")).Return("", nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On("CreateClass", mock.AnythingOfType("*model.Class")).
		Return(&model.Class{}, errors.New("insertion_error"))

	result, err := service.CreateClass(adminToken, mocks.MockCreateInput)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "insertion_error", err.Error())

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

func TestService_UpdateClass(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("UpdateClass", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.MockClass, nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On(
			"UpdateClass",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
			mock.AnythingOfType("model.Class")).
		Return(&mocks.MockUpdatedClass, nil)

	result, err := service.UpdateClass(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockUpdateInput)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_UpdateClass_CatchValidationErrors(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("UpdateClass", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.MockClass, nil)

	mockValidator.On("GetErrors").Return([]string{"validation_error"})

	result, err := service.UpdateClass(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockUpdateInput)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Validation errors: validation_error", err.Error())

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

func TestService_UpdateClass_CatchUpdateError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("UpdateClass", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.MockClass, nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.
		On("UpdateClass",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
			mock.AnythingOfType("model.Class")).
		Return(&model.Class{}, errors.New("update_error"))

	result, err := service.UpdateClass(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mocks.MockUpdateInput)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "update_error", err.Error())

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

func TestService_DeleteClassWithoutAdminToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("DeleteClass", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.MockClass, nil)

	mockRepo.
		On(
			"DeleteClass",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed", mock.AnythingOfType("model.Class")).
		Return(nil)

	err := service.DeleteClass(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)

	mockPolicy.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_DeleteClass_CatchDeleteError_WithoutAdminToken_AlreadySoftDeleted(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("DeleteClass", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.SoftDeletedMockClass, nil)

	err := service.DeleteClass(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.NotNil(t, err)
	assert.Equal(t, "class could not be deleted", err.Error())

	mockPolicy.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_DeleteClass_CatchDeleteError_WithAdminToken_AlreadySoftDeleted_NoFilter(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("DeleteClass", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.SoftDeletedMockClass, nil)

	err := service.DeleteClass(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.NotNil(t, err)
	assert.Equal(t, "class could not be deleted", err.Error())

	mockPolicy.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_GetClassByID(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("GetClass", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.MockClass, nil)

	result, err := service.GetClassById(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result.ID)
}

func TestService_ListClasses(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	//mockPolicy.On("ListClasses", mock.AnythingOfType("string")).Return(nil)

	mockPolicy.On("HasPermissions", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(true)

	mockRepo.On("ListClasses", mock.AnythingOfType("primitive.D"), mock.AnythingOfType("*options.FindOptions")).
		Return([]*model.ClassInfo{&mocks.MockClassInfo}, nil)

	mockValidator.On("GetErrors").Return([]string{})

	filter := model.ListClassFilter{}
	paginate := model.Paginator{}
	result, err := service.ListClasses(adminToken, &filter, &paginate)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.IsType(t, &model.ClassInfo{}, result[0])
	assert.Equal(t, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed", result[0].ID)

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestService_ListClasses_CatchRetrieveError(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockPolicy)
	service := &service2.ClassService{Validator: mockValidator, Repo: mockRepo, Policy: mockPolicy}

	mockPolicy.On("ListClasses", mock.AnythingOfType("string")).Return(nil)

	mockPolicy.On("HasPermissions", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(true)

	mockRepo.On("ListClasses", mock.AnythingOfType("primitive.D"), mock.AnythingOfType("*options.FindOptions")).Return([]*model.ClassInfo{}, errors.New("retrieval_error"))

	mockValidator.On("GetErrors").Return([]string{})

	filter := model.ListClassFilter{}
	paginate := model.Paginator{}
	result, err := service.ListClasses(adminToken, &filter, &paginate)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "retrieval_error", err.Error())

	mockPolicy.AssertExpectations(t)
}
