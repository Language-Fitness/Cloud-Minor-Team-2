package rpc

import (
	"ResultMicroservice/proto/result_pb"
	"ResultMicroservice/rpc"
	"ResultMicroservice/test/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestResultServer_DeleteByModule tests the DeleteByModule gRPC function.
func TestResultServer_DeleteByModule(t *testing.T) {
	mockResultService := new(mocks.MockResultService)
	resultServer := rpc.NewResultServer(mockResultService)
	request := &result_pb.DeleteByModuleRequest{
		BearerToken: "mockToken",
		ModuleID:    "mockModuleID",
	}
	mockResult := &result_pb.Response{
		ID:      "mockID",
		Deleted: true,
	}
	mockResultService.On("DeleteByModule", request.BearerToken, request.ModuleID).Return(mockResult.ID, mockResult.Deleted, nil)
	response, err := resultServer.DeleteByModule(context.Background(), request)
	mockResultService.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, mockResult, response)
}

// TestResultServer_SoftDeleteByModule tests the SoftDeleteByModule gRPC function.
func TestResultServer_SoftDeleteByModule(t *testing.T) {
	mockResultService := new(mocks.MockResultService)
	resultServer := rpc.NewResultServer(mockResultService)
	request := &result_pb.DeleteByModuleRequest{
		BearerToken: "mockToken",
		ModuleID:    "mockModuleID",
	}
	mockResult := &result_pb.Response{
		ID:      "mockID",
		Deleted: true,
	}
	mockResultService.On("SoftDeleteByModule", request.BearerToken, request.ModuleID).Return(mockResult.ID, mockResult.Deleted, nil)
	response, err := resultServer.SoftDeleteByModule(context.Background(), request)
	mockResultService.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, mockResult, response)
}

// TestResultServer_DeleteByClass tests the DeleteByClass gRPC function.
func TestResultServer_DeleteByClass(t *testing.T) {
	mockResultService := new(mocks.MockResultService)
	resultServer := rpc.NewResultServer(mockResultService)
	request := &result_pb.DeleteByClassRequest{
		BearerToken: "mockToken",
		ClassID:     "mockClassID",
	}
	mockResult := &result_pb.Response{
		ID:      "mockID",
		Deleted: true,
	}
	mockResultService.On("DeleteByClass", request.BearerToken, request.ClassID).Return(mockResult.ID, mockResult.Deleted, nil)
	response, err := resultServer.DeleteByClass(context.Background(), request)
	mockResultService.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, mockResult, response)
}

// TestResultServer_SoftDeleteByClass tests the SoftDeleteByClass gRPC function.
func TestResultServer_SoftDeleteByClass(t *testing.T) {
	mockResultService := new(mocks.MockResultService)
	resultServer := rpc.NewResultServer(mockResultService)
	request := &result_pb.DeleteByClassRequest{
		BearerToken: "mockToken",
		ClassID:     "mockClassID",
	}
	mockResult := &result_pb.Response{
		ID:      "mockID",
		Deleted: true,
	}
	mockResultService.On("SoftDeleteByClass", request.BearerToken, request.ClassID).Return(mockResult.ID, mockResult.Deleted, nil)
	response, err := resultServer.SoftDeleteByClass(context.Background(), request)
	mockResultService.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, mockResult, response)
}

// TestResultServer_DeleteByUser tests the DeleteByUser gRPC function.
func TestResultServer_DeleteByUser(t *testing.T) {
	mockResultService := new(mocks.MockResultService)
	resultServer := rpc.NewResultServer(mockResultService)
	request := &result_pb.DeleteByUserRequest{
		BearerToken: "mockToken",
		UserID:      "mockUserID",
	}
	mockResult := &result_pb.Response{
		ID:      "mockID",
		Deleted: true,
	}
	mockResultService.On("DeleteByUser", request.BearerToken, request.UserID).Return(mockResult.ID, mockResult.Deleted, nil)
	response, err := resultServer.DeleteByUser(context.Background(), request)
	mockResultService.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, mockResult, response)
}

// TestResultServer_SoftDeleteByUser tests the SoftDeleteByUser gRPC function.
func TestResultServer_SoftDeleteByUser(t *testing.T) {
	mockResultService := new(mocks.MockResultService)
	resultServer := rpc.NewResultServer(mockResultService)
	request := &result_pb.DeleteByUserRequest{
		BearerToken: "mockToken",
		UserID:      "mockUserID",
	}
	mockResult := &result_pb.Response{
		ID:      "mockID",
		Deleted: true,
	}
	mockResultService.On("SoftDeleteByUser", request.BearerToken, request.UserID).Return(mockResult.ID, mockResult.Deleted, nil)
	response, err := resultServer.SoftDeleteByUser(context.Background(), request)
	mockResultService.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, mockResult, response)
}

// TestResultServer_DeleteByModule_Error tests the DeleteByModule gRPC function for error scenarios.
func TestResultServer_DeleteByModule_Error(t *testing.T) {
	mockResultService := new(mocks.MockResultService)
	resultServer := rpc.NewResultServer(mockResultService)
	request := &result_pb.DeleteByModuleRequest{
		BearerToken: "mockToken",
		ModuleID:    "mockModuleID",
	}
	mockError := errors.New("mock error")
	mockResultService.On("DeleteByModule", request.BearerToken, request.ModuleID).Return("mockModuleID", false, mockError)
	response, err := resultServer.DeleteByModule(context.Background(), request)
	mockResultService.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

// TestResultServer_SoftDeleteByModule_Error tests the SoftDeleteByModule gRPC function for error scenarios.
func TestResultServer_SoftDeleteByModule_Error(t *testing.T) {
	mockResultService := new(mocks.MockResultService)
	resultServer := rpc.NewResultServer(mockResultService)
	request := &result_pb.DeleteByModuleRequest{
		BearerToken: "mockToken",
		ModuleID:    "mockModuleID",
	}
	mockError := errors.New("mock error")
	mockResultService.On("SoftDeleteByModule", request.BearerToken, request.ModuleID).Return("mockModuleID", false, mockError)
	response, err := resultServer.SoftDeleteByModule(context.Background(), request)
	mockResultService.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

// TestResultServer_DeleteByClass_Error tests the DeleteByClass gRPC function for error scenarios.
func TestResultServer_DeleteByClass_Error(t *testing.T) {
	mockResultService := new(mocks.MockResultService)
	resultServer := rpc.NewResultServer(mockResultService)
	request := &result_pb.DeleteByClassRequest{
		BearerToken: "mockToken",
		ClassID:     "mockClassID",
	}
	mockError := errors.New("mock error")
	mockResultService.On("DeleteByClass", request.BearerToken, request.ClassID).Return("mockModuleID", false, mockError)
	response, err := resultServer.DeleteByClass(context.Background(), request)
	mockResultService.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

// TestResultServer_SoftDeleteByClass_Error tests the SoftDeleteByClass gRPC function for error scenarios.
func TestResultServer_SoftDeleteByClass_Error(t *testing.T) {
	mockResultService := new(mocks.MockResultService)
	resultServer := rpc.NewResultServer(mockResultService)
	request := &result_pb.DeleteByClassRequest{
		BearerToken: "mockToken",
		ClassID:     "mockClassID",
	}
	mockError := errors.New("mock error")
	mockResultService.On("SoftDeleteByClass", request.BearerToken, request.ClassID).Return("mockModuleID", false, mockError)
	response, err := resultServer.SoftDeleteByClass(context.Background(), request)
	mockResultService.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

// TestResultServer_DeleteByUser_Error tests the DeleteByUser gRPC function for error scenarios.
func TestResultServer_DeleteByUser_Error(t *testing.T) {
	mockResultService := new(mocks.MockResultService)
	resultServer := rpc.NewResultServer(mockResultService)
	request := &result_pb.DeleteByUserRequest{
		BearerToken: "mockToken",
		UserID:      "mockUserID",
	}
	mockError := errors.New("mock error")
	mockResultService.On("DeleteByUser", request.BearerToken, request.UserID).Return("mockModuleID", false, mockError)
	response, err := resultServer.DeleteByUser(context.Background(), request)
	mockResultService.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

// TestResultServer_SoftDeleteByUser_Error tests the SoftDeleteByUser gRPC function for error scenarios.
func TestResultServer_SoftDeleteByUser_Error(t *testing.T) {
	mockResultService := new(mocks.MockResultService)
	resultServer := rpc.NewResultServer(mockResultService)
	request := &result_pb.DeleteByUserRequest{
		BearerToken: "mockToken",
		UserID:      "mockUserID",
	}
	mockError := errors.New("mock error")
	mockResultService.On("SoftDeleteByUser", request.BearerToken, request.UserID).Return("mockModuleID", false, mockError)
	response, err := resultServer.SoftDeleteByUser(context.Background(), request)
	mockResultService.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, response)
}
