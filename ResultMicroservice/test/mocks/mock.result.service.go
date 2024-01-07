package mocks

import (
	"ResultMicroservice/graph/model"
	"github.com/stretchr/testify/mock"
)

type MockResultService struct {
	mock.Mock
}

func (m *MockResultService) CreateResult(token string, newResult model.InputResult) (*model.Result, error) {
	args := m.Called(token, newResult)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultService) UpdateResult(token string, id string, updateData model.InputResult) (*model.Result, error) {
	args := m.Called(token, id, updateData)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultService) DeleteResult(token string, id string) (*model.Result, error) {
	args := m.Called(token, id)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultService) GetResultById(token string, id string) (*model.Result, error) {
	args := m.Called(token, id)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultService) ListResults(token string, filter *model.ResultFilter, paginate *model.Paginator) ([]*model.Result, error) {
	args := m.Called(token, filter, paginate)
	return args.Get(0).([]*model.Result), args.Error(1)
}
