package mocks

import (
	"Module/graph/model"
	"Module/internal/auth"
	"Module/internal/repository"
	"github.com/stretchr/testify/mock"
)

type MockPolicy struct {
	mock.Mock
	Token            auth.IToken
	ModuleRepository repository.IModuleRepository
}

func (m *MockPolicy) CreateModule(bearerToken string) error {
	args := m.Called(bearerToken)
	return args.Error(0)
}

func (m *MockPolicy) UpdateModule(bearerToken string, id string) (*model.Module, error) {
	args := m.Called(bearerToken, id)
	return args.Get(0).(*model.Module), args.Error(1)
}

func (m *MockPolicy) DeleteModule(bearerToken string, id string) error {
	args := m.Called(bearerToken, id)
	return args.Error(0)
}

func (m *MockPolicy) GetModule(bearerToken string) error {
	args := m.Called(bearerToken)
	return args.Error(0)
}

func (m *MockPolicy) ListModules(bearerToken string) error {
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
