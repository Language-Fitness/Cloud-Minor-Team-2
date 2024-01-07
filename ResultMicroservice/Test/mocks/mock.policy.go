package mocks

import (
	"ResultMicroservice/graph/model"
	"ResultMicroservice/internal/auth"
	"ResultMicroservice/internal/repository"
	"github.com/stretchr/testify/mock"
)

type MockResultPolicy struct {
	mock.Mock
	Token            auth.IToken
	ResultRepository repository.IResultRepository
}

func (m *MockResultPolicy) CreateResult(bearerToken string) error {
	args := m.Called(bearerToken)
	return args.Error(0)
}

func (m *MockResultPolicy) UpdateResult(bearerToken string, id string) (*model.Result, error) {
	args := m.Called(bearerToken, id)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultPolicy) DeleteResult(bearerToken string, id string) error {
	args := m.Called(bearerToken, id)
	return args.Error(0)
}

func (m *MockResultPolicy) GetResultByID(bearerToken string, id string) error {
	args := m.Called(bearerToken, id)
	return args.Error(0)
}

func (m *MockResultPolicy) ListResult(bearerToken string) error {
	args := m.Called(bearerToken)
	return args.Error(0)
}
