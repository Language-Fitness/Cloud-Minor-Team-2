package service

//import (
//	"Module/graph/model"
//	service2 "Module/internal/service"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"go.mongodb.org/mongo-driver/mongo"
//	"testing"
//)
//
//type MockRepository struct {
//	mock.Mock
//	modules    []*model.Module
//	collection *mongo.Collection
//}
//
//func (m *MockRepository) CreateModule(newModule *model.Module) (*model.Module, error) {
//	module := &model.Module{
//		ID:   "test-module-id",
//		Name: "Test Module",
//	}
//	return module, nil
//}
//
//func (m *MockRepository) UpdateModule(id string, updatedModule model.ModuleInput) (*model.Module, error) {
//	return nil, nil
//}
//
//func (m *MockRepository) DeleteModuleByID(id string) error {
//	return nil
//}
//
//func (m *MockRepository) GetModuleByID(id string) (*model.Module, error) {
//	return nil, nil
//}
//
//func (m *MockRepository) ListModules() ([]*model.Module, error) {
//	return nil, nil
//}
//
//type MockValidator struct {
//	mock.Mock
//}
//
//func (m *MockValidator) Validate(input interface{}, arr []string) {
//
//}
//
//func (m *MockValidator) GetErrors() []string {
//	args := m.Called()
//	return args.Get(0).([]string)
//}
//
//func TestService_CreateModule(t *testing.T) {
//	service := service2.NewModuleService()
//	mockRepo := new(MockRepository)
//	mockValidator := new(MockValidator)
//
//	newModule := model.ModuleInput{
//		Name: "Test Module",
//	}
//
//	//mockValidator.On("GetErrors").Return([]string{})
//	//mockRepo.On("CreateModule", mock.Anything).Return(module, nil)
//
//	service.Validator = mockValidator
//	service.Repo = mockRepo
//
//	// Call the CreateModule method and test the result
//	result, err := service.CreateModule(newModule)
//
//	// Assert that there's no error
//	assert.Nil(t, err)
//
//	// Add assertions to verify that the result matches your expectations
//	assert.NotNil(t, result)
//	assert.Equal(t, "test-module-id", result.ID)
//
//	// Assert that the expected functions were called on the mocks
//	mockValidator.AssertExpectations(t)
//	mockRepo.AssertExpectations(t)
//}
