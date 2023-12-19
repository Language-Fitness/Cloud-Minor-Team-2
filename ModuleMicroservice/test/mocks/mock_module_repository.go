package mocks

import (
	"Module/graph/model"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockRepository struct {
	mock.Mock
	modules    []*model.Module
	collection *mongo.Collection
}

func (m *MockRepository) CreateModule(newModule *model.Module) (*model.Module, error) {
	args := m.Called(newModule)
	return args.Get(0).(*model.Module), args.Error(1)
}

func (m *MockRepository) UpdateModule(id string, updatedModule model.Module) (*model.Module, error) {
	args := m.Called(id, updatedModule)
	return args.Get(0).(*model.Module), args.Error(1)
}

func (m *MockRepository) DeleteModuleByID(id string, existingClass model.Module) error {
	args := m.Called(id, existingClass)
	return args.Error(0)
}

func (m *MockRepository) GetModuleByID(id string) (*model.Module, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Module), args.Error(1)
}

func (m *MockRepository) ListModules(bsonFilter bson.D, paginateOptions *options.FindOptions) ([]*model.ModuleInfo, error) {
	args := m.Called(bsonFilter, paginateOptions)
	return args.Get(0).([]*model.ModuleInfo), args.Error(1)
}
