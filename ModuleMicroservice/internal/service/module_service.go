package service

import (
	"Module/graph/model"
	"Module/internal/auth"
	"Module/internal/repository"
	"Module/internal/validation"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

// IModuleService GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on Module.
type IModuleService interface {
	CreateModule(token string, newModule model.ModuleInput) (*model.Module, error)
	UpdateModule(token string, id string, updateData model.ModuleInput) (*model.Module, error)
	DeleteModule(token string, id string, filter *model.Filter) error
	GetModuleById(token string, id string) (*model.Module, error)
	ListModules(token string, filter *model.Filter, paginate *model.Paginator) ([]*model.ModuleInfo, error)
}

// ModuleService GOLANG STRUCT
// Contains two interfaces for a Validator and a Repo.
type ModuleService struct {
	Validator validation.IValidator
	Repo      repository.IModuleRepository
	Policy    auth.IPolicy
}

// NewModuleService GOLANG FACTORY
// Returns a ModuleService implementing IModuleService.
func NewModuleService(collection *mongo.Collection) IModuleService {
	return &ModuleService{
		Validator: validation.NewValidator(),
		Repo:      repository.NewModuleRepository(collection),
		Policy:    auth.NewPolicy(collection),
	}
}

func (m *ModuleService) CreateModule(token string, newModule model.ModuleInput) (*model.Module, error) {
	sub, err := m.Policy.CreateModule(token)
	if err != nil {
		return nil, err
	}

	m.Validator.Validate(newModule.Name, []string{"IsString", "Length:<25"}, "Name")
	m.Validator.Validate(newModule.Description, []string{"IsString", "Length:<50"}, "Description")
	m.Validator.Validate(newModule.Difficulty, []string{"IsInt"}, "Difficulty")
	m.Validator.Validate(newModule.Category, []string{"IsString"}, "Category")
	m.Validator.Validate(newModule.Private, []string{"IsBoolean"}, "Private")
	if newModule.Private {
		m.Validator.Validate(*newModule.Key, []string{"IsString", "Length:<30"}, "Key")
	}

	validationErrors := m.Validator.GetErrors()

	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		m.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()
	softDeleted := false

	moduleToInsert := &model.Module{
		ID:          uuid.New().String(),
		Name:        newModule.Name,
		Description: newModule.Description,
		Difficulty:  newModule.Difficulty,
		Category:    newModule.Category,
		MadeBy:      sub,
		Private:     newModule.Private,
		CreatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	}

	if newModule.Private {
		moduleToInsert.Key = newModule.Key
	}

	result, err := m.Repo.CreateModule(moduleToInsert)
	if err != nil {
		return nil, err
	}

	m.Validator.ClearErrors()
	return result, nil
}

func (m *ModuleService) UpdateModule(token string, id string, updateData model.ModuleInput) (*model.Module, error) {
	existingModule, err := m.Policy.UpdateModule(token, id)
	if err != nil {
		return nil, err
	}

	m.Validator.Validate(updateData.Name, []string{"IsString", "Length:<25"}, "Name")
	m.Validator.Validate(updateData.Description, []string{"IsString", "Length:<50"}, "Description")
	m.Validator.Validate(updateData.Difficulty, []string{"IsInt"}, "Difficulty")
	m.Validator.Validate(updateData.Category, []string{"IsString"}, "Category")
	m.Validator.Validate(updateData.Private, []string{"IsBoolean"}, "Private")
	if updateData.Private {
		m.Validator.Validate(*updateData.Key, []string{"IsString", "Length:<30"}, "Key")
	}

	validationErrors := m.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		m.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()
	newModule := model.Module{
		ID:          existingModule.ID,
		Name:        updateData.Name,
		Description: updateData.Description,
		Difficulty:  updateData.Difficulty,
		Category:    updateData.Category,
		MadeBy:      existingModule.MadeBy,
		Private:     updateData.Private,
		CreatedAt:   existingModule.CreatedAt,
		UpdatedAt:   &timestamp,
		SoftDeleted: existingModule.SoftDeleted,
	}

	if updateData.Private {
		newModule.Key = updateData.Key
	}

	result, err := m.Repo.UpdateModule(id, newModule)
	if err != nil {
		return nil, err
	}

	m.Validator.ClearErrors()
	return result, nil
}

func (m *ModuleService) DeleteModule(token string, id string, filter *model.Filter) error {
	//isAdmin, existingModule, err := m.Policy.DeleteModule(token, id)
	//if err != nil {
	//	return err
	//}

	//if !*existingModule.SoftDeleted {
	//	softDelete := true
	//	existingModule.SoftDeleted = &softDelete
	//
	//	err := m.Repo.SoftDeleteModuleByID(id, *existingModule)
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//}

	//
	//if isAdmin && filter != nil && !*filter.SoftDelete {
	//	err := m.Repo.HardDeleteModuleByID(id)
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//}

	return errors.New("module could not be deleted")
}

func (m *ModuleService) GetModuleById(token string, id string) (*model.Module, error) {
	existingModule, err := m.Policy.GetModule(token, id)
	if err != nil {
		return nil, err
	}

	return existingModule, nil
}

func (m *ModuleService) ListModules(token string, filter *model.Filter, paginate *model.Paginator) ([]*model.ModuleInfo, error) {
	//isAdmin, err := m.Policy.ListModules(token)
	//if err != nil {
	//	return nil, err
	//}

	m.Validator.Validate(filter.SoftDelete, []string{"IsNull", "IsBoolean"}, "Filter softDelete")
	m.Validator.Validate(filter.Name, []string{"IsNull", "IsString", "Length:<50"}, "Filter Name")
	m.Validator.Validate(filter.Description, []string{"IsNull", "IsString"}, "Filter Description")
	m.Validator.Validate(filter.Difficulty, []string{"IsNull", "IsInt"}, "Filter Difficulty")
	m.Validator.Validate(filter.Private, []string{"IsNull", "IsBoolean"}, "Filter Private")

	validationErrors := m.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		m.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	fmt.Println("test")
	fmt.Println(filter)
	fmt.Println(paginate)
	//fmt.Println(isAdmin)

	modules, err := m.Repo.ListModules()
	if err != nil {
		return nil, err
	}

	return modules, nil
}
