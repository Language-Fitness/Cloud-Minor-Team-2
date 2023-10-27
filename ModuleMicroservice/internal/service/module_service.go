package service

import (
	"Module/graph"
	"Module/graph/model"
	"Module/internal/repository"
	"Module/internal/validation"
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

// IModuleService GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on Module.
type IModuleService interface {
	CreateModule(newModule model.ModuleInput) (*model.Module, error)
	UpdateModule(id string, updatedModule model.ModuleInput) (*model.Module, error)
	DeleteModule(id string) error
	GetModuleById(id string) (*model.Module, error)
	ListModules() ([]*model.Module, error)
}

// ModuleService GOLANG STRUCT
// Contains two interfaces for a Validator and a Repo.
type ModuleService struct {
	Validator validation.IValidator
	Repo      repository.IModuleRepository
}

// NewModuleService GOLANG FACTORY
// Returns a ModuleService implementing IModuleService.
func NewModuleService(config *graph.AppConfig) IModuleService {
	return &ModuleService{
		Validator: validation.NewValidator(),
		Repo:      repository.NewModuleRepository(config.Collection),
	}
}

func (m *ModuleService) CreateModule(newModule model.ModuleInput) (*model.Module, error) {
	m.ValidateModuleInput(newModule)

	validationErrors := m.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
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
		MadeBy:      newModule.MadeBy,
		Private:     newModule.Private,
		CreatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	}
	result, err := m.Repo.CreateModule(moduleToInsert)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *ModuleService) UpdateModule(id string, updatedModule model.ModuleInput) (*model.Module, error) {
	m.ValidateModuleInput(updatedModule)

	validationErrors := m.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return nil, errors.New(errorMessage)
	}

	result, err := m.Repo.UpdateModule(id, updatedModule)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *ModuleService) DeleteModule(id string) error {
	m.Validator.Validate(id, []string{"IsUUID"})

	validationErrors := m.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return errors.New(errorMessage)
	}

	err := m.Repo.DeleteModuleByID(id)
	if err != nil {
		return err
	}

	return nil
}

func (m *ModuleService) GetModuleById(id string) (*model.Module, error) {
	m.Validator.Validate(id, []string{"IsUUID"})

	validationErrors := m.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return nil, errors.New(errorMessage)
	}

	module, err := m.Repo.GetModuleByID(id)
	if err != nil {
		return nil, err
	}

	return module, nil
}

func (m *ModuleService) ListModules() ([]*model.Module, error) {
	modules, err := m.Repo.ListModules()

	if err != nil {
		return nil, err
	}

	return modules, nil
}

func (m *ModuleService) ValidateModuleInput(module model.ModuleInput) {
	m.Validator.Validate(module.Name, []string{"IsString", "Length:<25"})
	m.Validator.Validate(*module.Description, []string{"IsString", "Length:<50"})
	m.Validator.Validate(*module.Difficulty, []string{"IsInt"})
	m.Validator.Validate(*module.Category, []string{"IsString"})
	m.Validator.Validate(*module.MadeBy, []string{"IsUUID"})
	m.Validator.Validate(*module.Private, []string{"IsBoolean"})
	m.Validator.Validate(*module.Key, []string{"IsString", "Length:<30"})
}
