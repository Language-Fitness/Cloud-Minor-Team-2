package service

import (
	"Module/graph/model"
	"Module/internal/database"
	"Module/internal/repository"
	"Module/internal/validation"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

// IModuleService GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on Module.
type IModuleService interface {
	CreateModule(newModule model.ModuleInput) (*model.Module, error)
	UpdateModule(id string, updateData model.ModuleInput) (*model.Module, error)
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
	m.Validator.Validate(newModule.Name, []string{"IsString", "Length:<25"})
	m.Validator.Validate(*newModule.Description, []string{"IsString", "Length:<50"})
	m.Validator.Validate(*newModule.Difficulty, []string{"IsInt"})
	m.Validator.Validate(*newModule.Category, []string{"IsString"})
	m.Validator.Validate(*newModule.MadeBy, []string{"IsUUID"})
	m.Validator.Validate(*newModule.Private, []string{"IsBoolean"})
	m.Validator.Validate(*newModule.Key, []string{"IsString", "Length:<30"})

	validationErrors := m.Validator.GetErrors()
	fmt.Println(validationErrors)
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

func (m *ModuleService) UpdateModule(id string, updateData model.ModuleInput) (*model.Module, error) {
	m.Validator.Validate(updateData.Name, []string{"IsString", "Length:<25"})
	m.Validator.Validate(*updateData.Description, []string{"IsString", "Length:<50"})
	m.Validator.Validate(*updateData.Difficulty, []string{"IsInt"})
	m.Validator.Validate(*updateData.Category, []string{"IsString"})
	m.Validator.Validate(*updateData.MadeBy, []string{"IsUUID"})
	m.Validator.Validate(*updateData.Private, []string{"IsBoolean"})
	m.Validator.Validate(*updateData.Key, []string{"IsString", "Length:<30"})

	validationErrors := m.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		m.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	existingModule, err := m.Repo.GetModuleByID(id)
	if err != nil {
		return nil, err
	}

	timestamp := time.Now().String()
	newModule := model.Module{
		ID:          existingModule.ID,
		Name:        updateData.Name,
		Description: updateData.Description,
		Difficulty:  updateData.Difficulty,
		Category:    updateData.Category,
		MadeBy:      updateData.MadeBy,
		Private:     updateData.Private,
		Key:         updateData.Key,
		CreatedAt:   existingModule.CreatedAt,
		UpdatedAt:   &timestamp,
		SoftDeleted: existingModule.SoftDeleted,
	}

	result, err := m.Repo.UpdateModule(id, newModule)
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
		m.Validator.ClearErrors()
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
		m.Validator.ClearErrors()
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
