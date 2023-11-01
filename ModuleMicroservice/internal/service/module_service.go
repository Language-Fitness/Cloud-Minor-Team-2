package service

import (
	"Module/graph/model"
	"Module/internal/database"
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
	UpdateModule(id string, updatedModule model.Module) (*model.Module, error)
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
func NewModuleService() IModuleService {
	collection, _ := database.GetCollection()

	return &ModuleService{
		Validator: validation.NewValidator(),
		Repo:      repository.NewModuleRepository(collection),
	}
}

func (m *ModuleService) CreateModule(newModule model.ModuleInput) (*model.Module, error) {
	m.ValidateModule(
		newModule.Name,
		newModule.Description,
		newModule.Difficulty,
		newModule.Category,
		newModule.MadeBy,
		newModule.Private,
		newModule.Key,
	)

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

	m.Validator.ClearErrors()
	return result, nil
}

func (m *ModuleService) UpdateModule(id string, updatedModule model.Module) (*model.Module, error) {
	m.ValidateModule(
		updatedModule.Name,
		updatedModule.Description,
		updatedModule.Difficulty,
		updatedModule.Category,
		updatedModule.MadeBy,
		updatedModule.Private,
		updatedModule.Key,
	)

	validationErrors := m.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()
	updatedModule.UpdatedAt = &timestamp

	result, err := m.Repo.UpdateModule(id, updatedModule)
	if err != nil {
		return nil, err
	}

	m.Validator.ClearErrors()
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

	m.Validator.ClearErrors()
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

	m.Validator.ClearErrors()
	return module, nil
}

func (m *ModuleService) ListModules() ([]*model.Module, error) {
	modules, err := m.Repo.ListModules()
	if err != nil {
		return nil, err
	}

	return modules, nil
}

func (m *ModuleService) ValidateModule(
	name string,
	description *string,
	difficulty *int,
	category *string,
	madeBy *string,
	private *bool,
	key *string,
) {
	m.Validator.Validate(name, []string{"IsString", "Length:<25"})
	m.Validator.Validate(description, []string{"IsString", "Length:<50"})
	m.Validator.Validate(difficulty, []string{"IsInt"})
	m.Validator.Validate(category, []string{"IsString"})
	m.Validator.Validate(madeBy, []string{"IsUUID"})
	m.Validator.Validate(private, []string{"IsBoolean"})
	m.Validator.Validate(key, []string{"IsString", "Length:<30"})
}
