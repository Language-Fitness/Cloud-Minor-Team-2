package service

import (
	"errors"
	"example/graph/model"
	"example/internal/database"
	"example/internal/repository"
	"example/internal/validation"
	"github.com/google/uuid"
	"strings"
	"time"
)

// IClassService GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on Class.
type IClassService interface {
	CreateClass(newClass model.ClassInput) (*model.Class, error)
	UpdateClass(id string, updatedClass model.Class) (*model.Class, error)
	DeleteClass(id string) error
	GetClassById(id string) (*model.Class, error)
	ListClasses() ([]*model.Class, error)
}

// ClassService GOLANG STRUCT
// Contains two interfaces for a Validator and a Repo.
type ClassService struct {
	Validator validation.IValidator
	Repo      repository.IClassRepository
}

// NewClassService GOLANG FACTORY
// Returns a ClassService implementing IClassService.
func NewClassService() IClassService {
	collection, _ := database.GetCollection()

	return &ClassService{
		Validator: validation.NewValidator(),
		Repo:      repository.NewClassRepository(collection),
	}
}

func (c *ClassService) CreateClass(newClass model.ClassInput) (*model.Class, error) {
	c.ValidateClass(
		newClass.ModuleID,
		newClass.Name,
		newClass.Description,
		newClass.Difficulty,
	)

	validationErrors := c.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()
	softDeleted := false

	ClassToInsert := &model.Class{
		ID:          uuid.New().String(),
		ModuleID:    newClass.ModuleID,
		Name:        newClass.Name,
		Description: newClass.Description,
		Difficulty:  newClass.Difficulty,
		CreatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	}

	result, err := c.Repo.CreateClass(ClassToInsert)
	if err != nil {
		return nil, err
	}

	c.Validator.ClearErrors()
	return result, nil
}

func (c *ClassService) UpdateClass(id string, updatedClass model.Class) (*model.Class, error) {
	c.ValidateClass(
		updatedClass.ModuleID,
		updatedClass.Name,
		updatedClass.Description,
		updatedClass.Difficulty,
	)

	validationErrors := c.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()
	updatedClass.UpdatedAt = &timestamp

	result, err := c.Repo.UpdateClass(id, updatedClass)
	if err != nil {
		return nil, err
	}

	c.Validator.ClearErrors()
	return result, nil
}

func (c *ClassService) DeleteClass(id string) error {
	c.Validator.Validate(id, []string{"IsUUID"})

	validationErrors := c.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return errors.New(errorMessage)
	}

	err := c.Repo.DeleteClassByID(id)
	if err != nil {
		return err
	}

	c.Validator.ClearErrors()
	return nil
}

func (c *ClassService) GetClassById(id string) (*model.Class, error) {
	c.Validator.Validate(id, []string{"IsUUID"})

	validationErrors := c.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return nil, errors.New(errorMessage)
	}

	Class, err := c.Repo.GetClassByID(id)
	if err != nil {
		return nil, err
	}

	c.Validator.ClearErrors()
	return Class, nil
}

func (c *ClassService) ListClasses() ([]*model.Class, error) {
	classes, err := c.Repo.ListClasses()
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (c *ClassService) ValidateClass(
	classId string,
	name string,
	description *string,
	difficulty *int,
) {
	c.Validator.Validate(classId, []string{"isUUID"})
	c.Validator.Validate(name, []string{"IsString", "Length:<25"})
	c.Validator.Validate(description, []string{"IsString", "Length:<50"})
	c.Validator.Validate(difficulty, []string{"IsInt"})
}
