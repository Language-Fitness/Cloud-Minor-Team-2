package mocks

import (
	"ResultMicroservice/graph/model"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockResultRepository struct {
	mock.Mock
	results    []*model.Result
	collection *mongo.Collection
}

func (m *MockResultRepository) CreateResult(newResult *model.Result) (*model.Result, error) {
	args := m.Called(newResult)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultRepository) UpdateResult(id string, updatedResult model.Result) (*model.Result, error) {
	args := m.Called(id, updatedResult)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultRepository) GetResultByID(id string) (*model.Result, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultRepository) ListResults(bsonFilter bson.D, paginateOptions *options.FindOptions) ([]*model.Result, error) {
	args := m.Called(bsonFilter, paginateOptions)
	return args.Get(0).([]*model.Result), args.Error(1)
}
