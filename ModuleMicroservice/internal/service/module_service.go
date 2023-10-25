package service

import (
	"Module/graph/model"
	"Module/internal/repository"
	"Module/internal/validation"
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type IModuleService interface {
	CreateModule(newModule *model.Module) error
	UpdateModule(updatedModule *model.Module) error
	DeleteModule(id string) error
	GetModuleById(id string) (*model.Module, error)
	ListModules() ([]*model.Module, error)
}

type ModuleService struct {
	validator *validation.Validator
	repo      *repository.ModuleRepository
}

func NewModuleService() *ModuleService {
	return &ModuleService{
		validator: validation.NewValidator(),
		repo:      repository.NewModuleRepository(),
	}
}

func (m *ModuleService) CreateModule(newModule model.ModuleInput) (*model.Module, error) {

	m.ValidateModuleInput(newModule)

	validationErrors := m.validator.GetErrors()
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
	result, err := m.repo.CreateModule(moduleToInsert)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *ModuleService) UpdateModule(id string, updatedModule model.ModuleInput) (*model.Module, error) {

	m.ValidateModuleInput(updatedModule)

	validationErrors := m.validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return nil, errors.New(errorMessage)
	}

	result, err := m.repo.UpdateModule(id, updatedModule)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *ModuleService) DeleteModule(id string) error {

	m.validator.Validate(id, []string{"IsUUID"})

	validationErrors := m.validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return errors.New(errorMessage)
	}

	err := m.repo.DeleteModuleByID(id)
	if err != nil {
		return err
	}

	return nil
}

func (m *ModuleService) GetModuleById(id string) (*model.Module, error) {

	m.validator.Validate(id, []string{"IsUUID"})

	validationErrors := m.validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return nil, errors.New(errorMessage)
	}

	module, err := m.repo.GetModuleByID(id)
	if err != nil {
		return nil, err
	}

	return module, nil
}

func (m *ModuleService) ListModules() ([]*model.Module, error) {
	modules, err := m.repo.ListModules()

	if err != nil {
		return nil, err
	}

	return modules, nil
}

func (m *ModuleService) ValidateModuleInput(module model.ModuleInput) {
	m.validator.Validate(module.Name, []string{"IsString", "Length:<25"})
	m.validator.Validate(*module.Description, []string{"IsString", "Length:<50"})
	m.validator.Validate(*module.Difficulty, []string{"IsInt"})
	m.validator.Validate(*module.Category, []string{"IsString"})
	m.validator.Validate(*module.MadeBy, []string{"IsString", "Length:<25"})
	m.validator.Validate(*module.Private, []string{"IsBoolean"})
	m.validator.Validate(*module.Key, []string{"IsString", "Length:<30"})
}
