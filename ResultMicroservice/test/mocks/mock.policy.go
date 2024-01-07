package mocks

import (
	"ResultMicroservice/graph/model"
	"github.com/stretchr/testify/mock"
)

type MockResultPolicy struct {
	mock.Mock
}

func (m *MockResultPolicy) CreateResult(bearerToken string) (string, error) {
	args := m.Called(bearerToken)
	return args.String(0), args.Error(1)
}

func (m *MockResultPolicy) UpdateResult(bearerToken string, id string) (*model.Result, error) {
	args := m.Called(bearerToken, id)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultPolicy) DeleteResult(bearerToken string, id string) (*model.Result, error) {
	args := m.Called(bearerToken, id)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultPolicy) GetResultByID(bearerToken string, id string) (*model.Result, error) {
	args := m.Called(bearerToken, id)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultPolicy) ListResult(bearerToken string) (string, bool, error) {
	args := m.Called(bearerToken)
	return args.String(0), args.Bool(1), args.Error(2)
}

func (m *MockResultPolicy) HasPermissions(bearerToken string, role string) bool {
	args := m.Called(bearerToken, role)
	return args.Bool(0)
}
