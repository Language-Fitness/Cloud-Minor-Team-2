package mocks

import (
	"example/graph/model"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockRepository struct {
	mock.Mock
	classes    []*model.Class
	collection *mongo.Collection
}

func (m *MockRepository) CreateClass(newClass *model.Class) (*model.Class, error) {
	args := m.Called(newClass)
	return args.Get(0).(*model.Class), args.Error(1)
}

func (m *MockRepository) UpdateClass(id string, updatedClass model.Class) (*model.Class, error) {
	args := m.Called(id, updatedClass)
	return args.Get(0).(*model.Class), args.Error(1)
}

func (m *MockRepository) SoftDeleteClassByID(id string, existingClass model.Class) error {
	args := m.Called(id, existingClass)
	return args.Error(0)
}

func (m *MockRepository) HardDeleteClassByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepository) GetClassByID(id string) (*model.Class, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Class), args.Error(1)
}

func (m *MockRepository) ListClasses() ([]*model.ClassInfo, error) {
	args := m.Called()
	return args.Get(0).([]*model.ClassInfo), args.Error(1)
}
