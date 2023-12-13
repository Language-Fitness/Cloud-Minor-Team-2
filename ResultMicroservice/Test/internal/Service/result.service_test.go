package service

import (
	"ResultMicroservice/graph/model"
	Service "ResultMicroservice/internal/service"
	"ResultMicroservice/test/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var adminToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJnWWlqam1Zd3Z5a2t3WUNlZUtpVzV3amxVM215dmVoNTRZSHlVZFc5MUFzIn0.eyJleHAiOjE3MDExMDgyMzEsImlhdCI6MTcwMTEwNzkzMSwianRpIjoiNzYxYTQwZjktNTMzMS00Mzc4LWI5OTktZjhjNWM3MGRkYWEzIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4ODg4L3JlYWxtcy9jbG91ZC1wcm9qZWN0IiwiYXVkIjpbInVzZXItbWFuYWdlbWVudC1jbGllbnQiLCJhY2NvdW50Il0sInN1YiI6IjQwN2VjMjNkLWM2ZjQtNDhkYi05YjFlLWZhN2Q3MDBmMjg2NiIsInR5cCI6IkJlYXJlciIsImF6cCI6ImxvZ2luLWNsaWVudCIsInNlc3Npb25fc3RhdGUiOiIwMWZhNjNkZi0wNDJmLTRmNTMtYmYzZi03NDNkYjFjMmY0MjYiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtY2xvdWQtcHJvamVjdCIsIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbInVwZGF0ZV9zY2hvb2wiLCJnZXRfY2xhc3NlcyIsImdldF9leGVyY2lzZXMiLCJkZWxldGVfbW9kdWxlIiwiZ2V0X3NjaG9vbHMiLCJkZWxldGVfZXhlcmNpc2UiLCJ1cGRhdGVfZXhlcmNpc2UiLCJnZXRfZXhlcmNpc2UiLCJkZWxldGVfbW9kdWxlX2FsbCIsImNyZWF0ZV9leGVyY2lzZSIsImdldF9zY2hvb2wiLCJkZWxldGVfZXhlcmNpc2VfYWxsIiwidXBkYXRlX3NjaG9vbF9hbGwiLCJkZWxldGVfY2xhc3MiLCJjcmVhdGVfbW9kdWxlIiwiZ2V0X21vZHVsZSIsInVwZGF0ZV9leGVyY2lzZV9hbGwiLCJnZXRfbW9kdWxlcyIsImNyZWF0ZV9jbGFzcyIsImNyZWF0ZV9zY2hvb2wiLCJ1cGRhdGVfbW9kdWxlX2FsbCIsImRlbGV0ZV9zY2hvb2wiLCJ1cGRhdGVfY2xhc3NfYWxsIiwidXBkYXRlX21vZHVsZSIsImdldF9jbGFzcyIsImRlbGV0ZV9zY2hvb2xfYWxsIiwidXBkYXRlX2NsYXNzIiwiZGVsZXRlX2NsYXNzX2FsbCJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwic2lkIjoiMDFmYTYzZGYtMDQyZi00ZjUzLWJmM2YtNzQzZGIxYzJmNDI2IiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiY2hhZCBhZG1pbiIsInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluQGFkbWluLmNvbSIsImdpdmVuX25hbWUiOiJjaGFkIiwiZmFtaWx5X25hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIn0.F4WBE9C3Ct17v5broRGPO92YR-lt9CzLprCnrOe4jWIMcMYyjHSBoLC-oQ7GHSoe1MjXe02CWRP98IZqQ5TPhF7nCliYs5qhn2vZRtlLa-QsjrTF2kZ1F_uEdXVekhVSKIRRFwoH8y2KxkaR3SSQ4bXOtJe8UJQs1AvzHPPeVDmRgfQcCZDNwdQTGI9Sb-8-C_dLXmU6W2ORJN1GmKikn9in4IS2kZ6KEiW6qNqOOllNlSQMZdtLQXf8BlymGf6s8z9j1itpg4iVljKeV8X76A8EHy-xQ98ESB188OVOxFHYReT82xOp5pusRjvMf3K71t20jPcOtUj-GiTZpEy9pQ"

func TestResultService_CreateResult(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("CreateResult", mock.AnythingOfType("string")).Return(nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("CreateResult", mock.AnythingOfType("*model.Result")).Return(&mocks.MockResult, nil)

	result, err := service.CreateResult(adminToken, mocks.MockInputResult)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "sample_result_id", result.ID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_CreateResult_CatchValidationErrors(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("CreateResult", mock.AnythingOfType("string")).Return(nil)

	mockValidator.On("GetErrors").Return([]string{"validation_error"})

	result, err := service.CreateResult(adminToken, mocks.MockInputResult)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Validation errors: validation_error", err.Error())

	mockValidator.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_CreateResult_CatchInsertError(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("CreateResult", mock.AnythingOfType("string")).Return(nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("CreateResult", mock.AnythingOfType("*model.Result")).Return(&model.Result{}, errors.New("insertion_error"))

	result, err := service.CreateResult(adminToken, mocks.MockInputResult)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "insertion_error", err.Error())

	mockValidator.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_UpdateResult(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("UpdateResult", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.MockResult, nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("UpdateResult", "sample_result_id", mock.AnythingOfType("model.Result")).Return(&mocks.MockResult, nil)

	result, err := service.UpdateResult(adminToken, "sample_result_id", mocks.MockInputResult)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "sample_result_id", result.ID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_UpdateResult_CatchValidationErrors(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("UpdateResult", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.MockResult, nil)

	mockValidator.On("GetErrors").Return([]string{"validation_error"})

	result, err := service.UpdateResult(adminToken, "sample_result_id", mocks.MockInputResult)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Validation errors: validation_error", err.Error())

	mockValidator.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_UpdateResult_CatchUpdateError(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("UpdateResult", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.MockResult, nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("UpdateResult", "sample_result_id", mock.AnythingOfType("model.Result")).
		Return(&model.Result{}, errors.New("update_error"))

	result, err := service.UpdateResult(adminToken, "sample_result_id", mocks.MockInputResult)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "update_error", err.Error())

	mockValidator.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_DeleteResult(t *testing.T) {
	// Create mocks
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)

	service := &Service.ResultService{
		Validator:    mockValidator,
		Repo:         mockRepo,
		ResultPolicy: mockPolicy,
	}

	mockPolicy.On("DeleteResult", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil).Once()

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("DeleteResultByID", "sample_result_id").Return(nil).Once()

	err := service.DeleteResult(adminToken, "sample_result_id")

	assert.Nil(t, err)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_DeleteResult_CatchPolicyError(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("DeleteResult", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(errors.New("policy_error"))

	err := service.DeleteResult(adminToken, "sample_result_id")

	assert.NotNil(t, err)
	assert.Equal(t, "policy_error", err.Error())

	mockRepo.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_DeleteResult_CatchDeleteError(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("DeleteResult", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("DeleteResultByID", "sample_result_id").Return(errors.New("delete_error"))

	err := service.DeleteResult(adminToken, "sample_result_id")

	assert.NotNil(t, err)
	assert.Equal(t, "delete_error", err.Error())

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_GetResultById(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("GetResultByID", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("GetResultByID", mock.AnythingOfType("string")).Return(&mocks.MockResult, nil)

	result, err := service.GetResultById(adminToken, "sample_result_id")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "sample_result_id", result.ID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_SoftDeleteByUser(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("SoftDeleteByUser", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("SoftDeleteByUser", mock.AnythingOfType("string")).Return(nil)

	userID, success, err := service.SoftDeleteByUser(adminToken, "sample_user_id")

	assert.Nil(t, err)
	assert.True(t, success)
	assert.Equal(t, "sample_user_id", userID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_SoftDeleteByClass(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("SoftDeleteByClass", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("SoftDeleteByClass", mock.AnythingOfType("string")).Return(nil)

	classID, success, err := service.SoftDeleteByClass(adminToken, "sample_class_id")

	assert.Nil(t, err)
	assert.True(t, success)
	assert.Equal(t, "sample_class_id", classID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_SoftDeleteByModule(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("SoftDeleteByModule", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("SoftDeleteByModule", mock.AnythingOfType("string")).Return(nil)

	moduleID, success, err := service.SoftDeleteByModule(adminToken, "sample_module_id")

	assert.Nil(t, err)
	assert.True(t, success)
	assert.Equal(t, "sample_module_id", moduleID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_DeleteByUser(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("DeleteByUser", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("DeleteByUser", mock.AnythingOfType("string")).Return(nil)

	userID, success, err := service.DeleteByUser(adminToken, "sample_user_id")

	assert.Nil(t, err)
	assert.True(t, success)
	assert.Equal(t, "sample_user_id", userID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_DeleteByClass(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("DeleteByClass", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("DeleteByClass", mock.AnythingOfType("string")).Return(nil)

	classID, success, err := service.DeleteByClass(adminToken, "sample_class_id")

	assert.Nil(t, err)
	assert.True(t, success)
	assert.Equal(t, "sample_class_id", classID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}

func TestResultService_DeleteByModule(t *testing.T) {
	mockRepo := new(mocks.MockResultRepository)
	mockValidator := new(mocks.MockValidator)
	mockPolicy := new(mocks.MockResultPolicy)
	service := &Service.ResultService{Validator: mockValidator, Repo: mockRepo, ResultPolicy: mockPolicy}

	mockPolicy.On("DeleteByModule", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("DeleteByModule", mock.AnythingOfType("string")).Return(nil)

	moduleID, success, err := service.DeleteByModule(adminToken, "sample_module_id")

	assert.Nil(t, err)
	assert.True(t, success)
	assert.Equal(t, "sample_module_id", moduleID)

	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockPolicy.AssertExpectations(t)
}
