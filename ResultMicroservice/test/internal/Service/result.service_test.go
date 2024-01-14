package service

import (
	"ResultMicroservice/graph/model"
	"ResultMicroservice/internal/service"
	"ResultMicroservice/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var userToken = "user_token"

func TestResultService_CreateResult(t *testing.T) {
	mockValidator := new(mocks.MockValidator)
	mockRepo := new(mocks.MockResultRepository)
	mockPolicy := new(mocks.MockResultPolicy)

	resultService := service.ResultService{
		Validator:    mockValidator,
		Repo:         mockRepo,
		ResultPolicy: mockPolicy,
	}

	mockPolicy.On("CreateResult", mock.AnythingOfType("string")).Return("", nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("CreateResult", mock.AnythingOfType("*model.Result")).Return(&model.Result{ID: "some_result_id"}, nil)

	newResult := model.InputResult{
		ExerciseID: "some_exercise_id",
		ClassID:    "some_class_id",
		ModuleID:   "some_module_id",
		Input:      "some_input",
		Result:     true,
	}

	result, err := resultService.CreateResult(userToken, newResult)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "some_result_id", result.ID)

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestResultService_UpdateResult(t *testing.T) {
	mockValidator := new(mocks.MockValidator)
	mockRepo := new(mocks.MockResultRepository)
	mockPolicy := new(mocks.MockResultPolicy)

	resultService := service.ResultService{
		Validator:    mockValidator,
		Repo:         mockRepo,
		ResultPolicy: mockPolicy,
	}

	mockPolicy.On("UpdateResult", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&model.Result{}, nil)

	mockValidator.On("GetErrors").Return([]string{})

	mockRepo.On("UpdateResult", mock.AnythingOfType("string"), mock.AnythingOfType("model.Result")).
		Return(&model.Result{ID: "some_result_id"}, nil)

	updateData := model.InputResult{
		ExerciseID: "some_updated_exercise_id",
		ClassID:    "some_updated_class_id",
		ModuleID:   "some_updated_module_id",
		Input:      "some_updated_input",
		Result:     false,
	}

	result, err := resultService.UpdateResult(userToken, "some_result_id", updateData)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "some_result_id", result.ID)

	mockPolicy.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestResultService_DeleteResult(t *testing.T) {
	mockValidator := new(mocks.MockValidator)
	mockRepo := new(mocks.MockResultRepository)
	mockPolicy := new(mocks.MockResultPolicy)

	resultService := service.ResultService{
		Validator:    mockValidator,
		Repo:         mockRepo,
		ResultPolicy: mockPolicy,
	}

	mockValidator.On("GetErrors").Return([]string{}).Return([]string{})

	mockPolicy.On("DeleteResult", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&model.Result{}, nil)

	mockRepo.On("UpdateResult", mock.AnythingOfType("string"), mock.AnythingOfType("model.Result")).
		Return(&mocks.MockResult, nil)

	err := resultService.DeleteResult(userToken, "some_result_id", true)

	assert.Nil(t, err)

	mockPolicy.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

func TestResultService_UnDeleteResult(t *testing.T) {
	mockValidator := new(mocks.MockValidator)
	mockRepo := new(mocks.MockResultRepository)
	mockPolicy := new(mocks.MockResultPolicy)

	resultService := service.ResultService{
		Validator:    mockValidator,
		Repo:         mockRepo,
		ResultPolicy: mockPolicy,
	}

	mockValidator.On("GetErrors").Return([]string{})

	mockPolicy.On("DeleteResult", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&mocks.MockDeletedResult, nil)

	mockRepo.On("UpdateResult", mock.AnythingOfType("string"), mock.AnythingOfType("model.Result")).
		Return(&model.Result{ID: "some_result_id"}, nil)

	err := resultService.DeleteResult(userToken, "some_result_id", false)

	assert.Nil(t, err)

	mockPolicy.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

func TestResultService_GetResultById(t *testing.T) {
	mockValidator := new(mocks.MockValidator)
	mockRepo := new(mocks.MockResultRepository)
	mockPolicy := new(mocks.MockResultPolicy)

	resultService := service.ResultService{
		Validator:    mockValidator,
		Repo:         mockRepo,
		ResultPolicy: mockPolicy,
	}

	mockValidator.On("GetErrors").Return([]string{})

	mockPolicy.On("GetResultByID", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&model.Result{ID: "some_result_id"}, nil)

	result, err := resultService.GetResultById(userToken, "some_result_id")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "some_result_id", result.ID)

	mockPolicy.AssertExpectations(t)
}

func TestResultService_ListResults(t *testing.T) {
	mockValidator := new(mocks.MockValidator)
	mockRepo := new(mocks.MockResultRepository)
	mockPolicy := new(mocks.MockResultPolicy)

	resultService := service.ResultService{
		Validator:    mockValidator,
		Repo:         mockRepo,
		ResultPolicy: mockPolicy,
	}

	mockPolicy.On("ListResult", mock.AnythingOfType("string")).
		Return("some-id", true, nil)

	mockValidator.On("GetErrors").Return([]string{}).Return([]string{})

	mockRepo.On("ListResults", mock.AnythingOfType("primitive.D"), mock.AnythingOfType("*options.FindOptions")).
		Return([]*model.ResultInfo{&model.ResultInfo{ID: "some_result_id"}}, nil)

	filter := &model.ResultFilter{}
	paginate := &model.Paginator{}

	results, err := resultService.ListResults(userToken, filter, paginate)

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, "some_result_id", results[0].ID)

	mockPolicy.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}
