package service

import (
	"errors"
	"example/graph/model"
	"example/internal/auth"
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
	CreateClass(token string, newClass model.ClassInput) (*model.Class, error)
	UpdateClass(token string, id string, updatedData model.ClassInput) (*model.Class, error)
	DeleteClass(token string, id string, filter *model.Filter) error
	GetClassById(token string, id string) (*model.Class, error)
	ListClasses(token string) ([]*model.ClassInfo, error)
}

// ClassService GOLANG STRUCT
// Contains two interfaces for a Validator and a Repo.
type ClassService struct {
	Validator validation.IValidator
	Repo      repository.IClassRepository
	Policy    auth.IPolicy
}

// NewClassService GOLANG FACTORY
// Returns a ClassService implementing IClassService.
func NewClassService() IClassService {
	collection, _ := database.GetCollection()

	return &ClassService{
		Validator: validation.NewValidator(),
		Repo:      repository.NewClassRepository(collection),
		Policy:    auth.NewPolicy(collection),
	}
}

func (c *ClassService) CreateClass(token string, newClass model.ClassInput) (*model.Class, error) {
	sub, err := c.Policy.CreateClass(token)
	if err != nil {
		return nil, err
	}

	c.Validator.Validate(newClass.ModuleID, []string{"IsUUID"})
	c.Validator.Validate(newClass.Name, []string{"IsString", "Length:<25"})
	c.Validator.Validate(newClass.Description, []string{"IsString", "Length:<50"})
	c.Validator.Validate(newClass.Difficulty, []string{"IsInt"})

	validationErrors := c.Validator.GetErrors()

	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		c.Validator.ClearErrors()
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
		MadeBy:      sub,
		CreatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	}

	result, err := c.Repo.CreateClass(ClassToInsert)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ClassService) UpdateClass(token string, id string, updatedData model.ClassInput) (*model.Class, error) {
	existingClass, err := c.Policy.UpdateClass(token, id)
	if err != nil {
		return nil, err
	}

	c.Validator.Validate(updatedData.ModuleID, []string{"IsUUID"})
	c.Validator.Validate(updatedData.Name, []string{"IsString", "Length:<25"})
	c.Validator.Validate(updatedData.Description, []string{"IsString", "Length:<50"})
	c.Validator.Validate(updatedData.Difficulty, []string{"IsInt"})

	validationErrors := c.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		c.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()
	newClass := model.Class{
		ID:          existingClass.ID,
		ModuleID:    updatedData.ModuleID,
		Name:        updatedData.Name,
		Description: updatedData.Description,
		Difficulty:  updatedData.Difficulty,
		MadeBy:      existingClass.MadeBy,
		CreatedAt:   existingClass.CreatedAt,
		UpdatedAt:   &timestamp,
		SoftDeleted: existingClass.SoftDeleted,
	}

	result, err := c.Repo.UpdateClass(id, newClass)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ClassService) DeleteClass(token string, id string, filter *model.Filter) error {
	isAdmin, existingClass, err := c.Policy.DeleteClass(token, id)
	if err != nil {
		return err
	}

	if !*existingClass.SoftDeleted {
		softDelete := true
		existingClass.SoftDeleted = &softDelete

		err := c.Repo.SoftDeleteClassByID(id, *existingClass)
		if err != nil {
			return err
		}
		return nil
	}

	if isAdmin && filter != nil && !*filter.SoftDelete {
		err := c.Repo.HardDeleteClassByID(id)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("class could not be deleted")
}

func (c *ClassService) GetClassById(token string, id string) (*model.Class, error) {
	existingClass, err := c.Policy.GetClass(token, id)
	if err != nil {
		return nil, err
	}

	return existingClass, nil
}

func (c *ClassService) ListClasses(token string) ([]*model.ClassInfo, error) {
	err := c.Policy.ListClasses(token)
	if err != nil {
		return nil, err
	}

	classes, err := c.Repo.ListClasses()
	if err != nil {
		return nil, err
	}

	return classes, nil
}
