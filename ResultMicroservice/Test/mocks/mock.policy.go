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

func (m *MockResultPolicy) GetResultByExercise(bearerToken string, exerciseID string) error {
	args := m.Called(bearerToken, exerciseID)
	return args.Error(0)
}

func (m *MockResultPolicy) GetResultsByClass(bearerToken string, classID string) error {
	args := m.Called(bearerToken, classID)
	return args.Error(0)
}

func (m *MockResultPolicy) GetResultsByUser(bearerToken string, userID string) error {
	args := m.Called(bearerToken, userID)
	return args.Error(0)
}

func (m *MockResultPolicy) GetResultByID(bearerToken string, id string) error {
	args := m.Called(bearerToken, id)
	return args.Error(0)
}

func (m *MockResultPolicy) DeleteResultByClass(bearerToken string, classID string) (*string, error) {
	args := m.Called(bearerToken, classID)
	return args.Get(0).(*string), args.Error(1)
}

func (m *MockResultPolicy) SoftDeleteByUser(bearerToken string, userID string) error {
	args := m.Called(bearerToken, userID)
	return args.Error(0)
}

func (m *MockResultPolicy) SoftDeleteByClass(bearerToken string, classID string) error {
	args := m.Called(bearerToken, classID)
	return args.Error(0)
}

func (m *MockResultPolicy) SoftDeleteByModule(bearerToken string, moduleID string) error {
	args := m.Called(bearerToken, moduleID)
	return args.Error(0)
}

func (m *MockResultPolicy) DeleteByUser(bearerToken string, userID string) error {
	args := m.Called(bearerToken, userID)
	return args.Error(0)
}

func (m *MockResultPolicy) DeleteByClass(bearerToken string, classID string) error {
	args := m.Called(bearerToken, classID)
	return args.Error(0)
}

func (m *MockResultPolicy) DeleteByModule(bearerToken string, moduleID string) error {
	args := m.Called(bearerToken, moduleID)
	return args.Error(0)
}
