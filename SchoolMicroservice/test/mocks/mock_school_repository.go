package mocks

import (
	"example/graph/model"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockRepository struct {
	mock.Mock
	Schools    []*model.School
	collection *mongo.Collection
}

func (m *MockRepository) CreateSchool(newSchool *model.School) (*model.School, error) {
	args := m.Called(newSchool)
	return args.Get(0).(*model.School), args.Error(1)
}

func (m *MockRepository) UpdateSchool(id string, updatedSchool model.School) (*model.School, error) {
	args := m.Called(id, updatedSchool)
	return args.Get(0).(*model.School), args.Error(1)
}

func (m *MockRepository) DeleteSchool(id string, existingSchool model.School) error {
	args := m.Called(id, existingSchool)
	return args.Error(0)
}

func (m *MockRepository) GetSchoolByID(id string) (*model.School, error) {
	args := m.Called(id)
	return args.Get(0).(*model.School), args.Error(1)
}

func (m *MockRepository) ListSchools(bsonFilter bson.D, paginateOptions *options.FindOptions) ([]*model.SchoolInfo, error) {
	args := m.Called()
	return args.Get(0).([]*model.SchoolInfo), args.Error(1)
}
