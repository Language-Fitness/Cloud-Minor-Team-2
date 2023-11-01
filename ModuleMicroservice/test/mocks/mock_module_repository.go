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
	module := &model.Module{
		ID:   "test-module-id",
		Name: "Test Module",
	}
	return module, nil
}

func (m *MockRepository) UpdateModule(id string, updatedModule model.Module) (*model.Module, error) {
	return nil, nil
}

func (m *MockRepository) DeleteModuleByID(id string) error {
	return nil
}

func (m *MockRepository) GetModuleByID(id string) (*model.Module, error) {
	return nil, nil
}

func (m *MockRepository) ListModules() ([]*model.Module, error) {
	return nil, nil
}
