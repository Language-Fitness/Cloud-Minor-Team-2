package mocks

import (
	"Module/graph/model"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
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

func (m *MockRepository) DeleteModuleByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepository) GetModuleByID(id string) (*model.Module, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Module), args.Error(1)
}

func (m *MockRepository) ListModules() ([]*model.ModuleInfo, error) {
	args := m.Called()
	return args.Get(0).([]*model.ModuleInfo), args.Error(1)
}
