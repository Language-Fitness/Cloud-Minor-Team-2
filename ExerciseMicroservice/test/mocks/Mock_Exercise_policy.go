package mocks

import (
	"ExerciseMicroservice/graph/model"
	"github.com/stretchr/testify/mock"
)

type MockExercisePolicy struct {
	mock.Mock
}

func (m *MockExercisePolicy) CreateExercise(bearerToken string) (string, error) {
	args := m.Called(bearerToken)
	return args.String(0), args.Error(1)
}

func (m *MockExercisePolicy) UpdateExercise(bearerToken string, id string) (*model.Exercise, error) {
	args := m.Called(bearerToken, id)
	return args.Get(0).(*model.Exercise), args.Error(1)
}

func (m *MockExercisePolicy) DeleteExercise(bearerToken string, id string) (bool, *model.Exercise, error) {
	args := m.Called(bearerToken, id)
	return args.Bool(0), args.Get(1).(*model.Exercise), args.Error(2)
}

func (m *MockExercisePolicy) GetExercise(bearerToken string, id string) (*model.Exercise, error) {
	args := m.Called(bearerToken, id)
	return args.Get(0).(*model.Exercise), args.Error(1)
}

func (m *MockExercisePolicy) ListExercises(bearerToken string) (bool, error) {
	args := m.Called(bearerToken)
	return args.Bool(0), args.Error(1)
}

func (m *MockExercisePolicy) HasPermissions(bearerToken string, role string) bool {
	args := m.Called(bearerToken, role)
	return args.Bool(0)
}
