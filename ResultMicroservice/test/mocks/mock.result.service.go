package mocks

import (
	"ResultMicroservice/graph/model"
	"github.com/stretchr/testify/mock"
)

type MockResultService struct {
	mock.Mock
}

func (m *MockResultService) CreateResult(bearerToken string, newResult model.InputResult) (*model.Result, error) {
	args := m.Called(bearerToken, newResult)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultService) UpdateResult(bearerToken string, id string, updateData model.InputResult) (*model.Result, error) {
	args := m.Called(bearerToken, id, updateData)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultService) DeleteResult(bearerToken string, id string) error {
	args := m.Called(bearerToken, id)
	return args.Error(0)
}

func (m *MockResultService) GetResultById(bearerToken string, id string) (*model.Result, error) {
	args := m.Called(bearerToken, id)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultService) ListResults(bearerToken string) ([]*model.Result, error) {
	args := m.Called(bearerToken)
	return args.Get(0).([]*model.Result), args.Error(1)
}
