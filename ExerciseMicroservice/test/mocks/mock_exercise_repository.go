package mocks

import (
	"ExerciseMicroservice/graph/model"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockExerciseRepository struct {
	mock.Mock
	exercises  []*model.ExerciseInfo
	collection *mongo.Collection
}

func (m *MockExerciseRepository) CreateExercise(newExercise *model.Exercise) (*model.Exercise, error) {
	args := m.Called(newExercise)
	return args.Get(0).(*model.Exercise), args.Error(1)
}

func (m *MockExerciseRepository) UpdateExercise(id string, updatedExercise model.Exercise) (*model.Exercise, error) {
	args := m.Called(id, updatedExercise)
	return args.Get(0).(*model.Exercise), args.Error(1)
}

func (m *MockExerciseRepository) GetExerciseByID(id string) (*model.Exercise, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Exercise), args.Error(1)
}

func (m *MockExerciseRepository) ListExercises(bsonFilter bson.D, paginateOptions *options.FindOptions) ([]*model.ExerciseInfo, error) {
	args := m.Called(bsonFilter, paginateOptions)
	return args.Get(0).([]*model.ExerciseInfo), args.Error(1)
}
