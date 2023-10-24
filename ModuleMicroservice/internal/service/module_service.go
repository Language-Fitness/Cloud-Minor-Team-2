package service

import (
	"Module/graph/model"
	"Module/internal/repository"
	"Module/internal/validation"
	"errors"
	"strings"
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

func NewModuleService(repository *repository.ModuleRepository) *ModuleService {
	return &ModuleService{
		validator: validation.NewValidator(),
		repo:      repository,
	}
}

func (m *ModuleService) CreateModule(newModule *model.Module) error {

	m.ValidateModule(newModule)

	validationErrors := m.validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return errors.New(errorMessage)
	}

	err := m.repo.CreateModule(newModule)

	if err != nil {
		return err
	}

	return nil
}

func (m *ModuleService) UpdateModule(updatedModule *model.Module) error {

	m.ValidateModule(updatedModule)

	validationErrors := m.validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return errors.New(errorMessage)
	}

	err := m.repo.UpdateModule(updatedModule)

	if err != nil {
		return err
	}

	return nil
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

func (m *ModuleService) ValidateModule(module *model.Module) {
	m.validator.Validate(module.ID, []string{"IsUUID"})
	m.validator.Validate(module.Name, []string{"IsString", "Length:<25"})
	m.validator.Validate(*module.Description, []string{"IsString", "Length:<50"})
	m.validator.Validate(*module.Difficulty, []string{"IsInt"})
	m.validator.Validate(*module.Category, []string{"IsString"})
	m.validator.Validate(*module.MadeBy, []string{"IsString", "Length:<25"})
	m.validator.Validate(*module.Private, []string{"IsBoolean"})
	m.validator.Validate(*module.Key, []string{"IsString", "Length:<30"})
	m.validator.Validate(*module.CreatedAt, []string{"IsDatetime"})
	m.validator.Validate(*module.UpdatedAt, []string{"IsDatetime"})
	m.validator.Validate(*module.SoftDeleted, []string{"IsBoolean"})
}
