package mocks

import (
	"ResultMicroservice/graph/model"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
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

func (m *MockResultRepository) DeleteResultByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockResultRepository) GetResultByID(id string) (*model.Result, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Result), args.Error(1)
}

func (m *MockResultRepository) DeleteResultByClassAndUserID(classID string, userID string) error {
	args := m.Called(classID, userID)
	return args.Error(0)
}

func (m *MockResultRepository) SoftDeleteByUser(userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *MockResultRepository) SoftDeleteByClass(classID string) error {
	args := m.Called(classID)
	return args.Error(0)
}

func (m *MockResultRepository) SoftDeleteByModule(moduleID string) error {
	args := m.Called(moduleID)
	return args.Error(0)
}

func (m *MockResultRepository) DeleteByUser(userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *MockResultRepository) DeleteByClass(classID string) error {
	args := m.Called(classID)
	return args.Error(0)
}

func (m *MockResultRepository) DeleteByModule(moduleID string) error {
	args := m.Called(moduleID)
	return args.Error(0)
}

func (m *MockResultRepository) ListResults() ([]*model.Result, error) {
	args := m.Called()
	return args.Get(0).([]*model.Result), args.Error(1)
}
