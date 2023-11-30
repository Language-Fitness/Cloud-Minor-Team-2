package mocks

import (
	"example/graph/model"
	"example/internal/auth"
	"example/internal/repository"
	"github.com/stretchr/testify/mock"
)

type MockPolicy struct {
	mock.Mock
	Token            auth.IToken
	SchoolRepository repository.ISchoolRepository
}

func (m *MockPolicy) CreateSchool(bearerToken string) error {
	args := m.Called(bearerToken)
	return args.Error(0)
}

func (m *MockPolicy) UpdateSchool(bearerToken string, id string) (*model.School, error) {
	args := m.Called(bearerToken, id)
	return args.Get(0).(*model.School), args.Error(1)
}

func (m *MockPolicy) DeleteSchool(bearerToken string, id string) error {
	args := m.Called(bearerToken, id)
	return args.Error(0)
}

func (m *MockPolicy) GetSchool(bearerToken string) error {
	args := m.Called(bearerToken)
	return args.Error(0)
}

func (m *MockPolicy) ListSchools(bearerToken string) error {
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
