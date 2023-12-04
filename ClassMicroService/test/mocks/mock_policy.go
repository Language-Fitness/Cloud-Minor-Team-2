package mocks

import (
	"example/graph/model"
	"example/internal/auth"
	"example/internal/repository"
	"github.com/stretchr/testify/mock"
)

type MockPolicy struct {
	mock.Mock
	Token           auth.IToken
	ClassRepository repository.IClassRepository
}

func (m *MockPolicy) CreateClass(bearerToken string) (string, error) {
	args := m.Called(bearerToken)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockPolicy) UpdateClass(bearerToken string, id string) (*model.Class, error) {
	args := m.Called(bearerToken, id)
	return args.Get(0).(*model.Class), args.Error(1)
}

func (m *MockPolicy) DeleteClass(bearerToken string, id string) (bool, *model.Class, error) {
	args := m.Called(bearerToken, id)
	return args.Get(0).(bool), args.Get(1).(*model.Class), args.Error(2)
}

func (m *MockPolicy) GetClass(bearerToken string, id string) (*model.Class, error) {
	args := m.Called(bearerToken, id)
	return args.Get(0).(*model.Class), args.Error(1)
}

func (m *MockPolicy) ListClasses(bearerToken string) error {
	args := m.Called(bearerToken)
	return args.Error(0)
}

func (m *MockPolicy) getSubAndRoles(bearerToken string) (string, []interface{}, error) {
	args := m.Called(bearerToken)
	return args.String(0), args.Get(1).([]interface{}), args.Error(2)
}

func (m *MockPolicy) hasRole(roles []interface{}, targetRole string) bool {
	args := m.Called(roles, targetRole)
	return args.Bool(0)
}
