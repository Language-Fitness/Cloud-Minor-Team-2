package mocks

import (
	"ExerciseMicroservice/graph/model"
	"github.com/stretchr/testify/mock"
)

type MockExerciseService struct {
	mock.Mock
}

func (m *MockExerciseService) CreateExercise(token string, newExercise model.ExerciseInput) (*model.Exercise, error) {
	args := m.Called(token, newExercise)
	return args.Get(0).(*model.Exercise), args.Error(1)
}

func (m *MockExerciseService) UpdateExercise(token string, id string, updateData model.ExerciseInput) (*model.Exercise, error) {
	args := m.Called(token, id, updateData)
	return args.Get(0).(*model.Exercise), args.Error(1)
}

func (m *MockExerciseService) DeleteExercise(token string, id string, deleteFlag bool) error {
	args := m.Called(token, id, deleteFlag)
	return args.Error(0)
}

func (m *MockExerciseService) UnDeleteExercise(token string, id string) error {
	args := m.Called(token, id)
	return args.Error(0)
}

func (m *MockExerciseService) GetExerciseById(token string, id string) (*model.Exercise, error) {
	args := m.Called(token, id)
	return args.Get(0).(*model.Exercise), args.Error(1)
}

func (m *MockExerciseService) ListExercises(token string, filter *model.ExerciseFilter, paginate *model.Paginator) ([]*model.ExerciseInfo, error) {
	args := m.Called(token, filter, paginate)
	return args.Get(0).([]*model.ExerciseInfo), args.Error(1)
}
