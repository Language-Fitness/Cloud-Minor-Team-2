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

// ISchoolService GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on School.
type ISchoolService interface {
	CreateSchool(newSchool model.SchoolInput) (*model.School, error)
	UpdateSchool(id string, updatedData model.SchoolInput) (*model.School, error)
	DeleteSchool(id string) error
	GetSchoolById(id string) (*model.School, error)
	ListSchools() ([]*model.School, error)
}

// SchoolService GOLANG STRUCT
// Contains two interfaces for a Validator and a Repo.
type SchoolService struct {
	Validator validation.IValidator
	Repo      repository.ISchoolRepository
}

// NewSchoolService GOLANG FACTORY
// Returns a SchoolService implementing ISchoolService.
func NewSchoolService() ISchoolService {
	collection, _ := database.GetCollection()

	return &SchoolService{
		Validator: validation.NewValidator(),
		Repo:      repository.NewSchoolRepository(collection),
	}
}

func (c *SchoolService) CreateSchool(newSchool model.SchoolInput) (*model.School, error) {
	c.Validator.Validate(newSchool.Name, []string{"IsString", "Length:<25"})
	c.Validator.Validate(*newSchool.Location, []string{"IsString", "Length:<50"})

	validationErrors := c.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()
	softDeleted := false

	SchoolToInsert := &model.School{
		ID:          uuid.New().String(),
		Name:        newSchool.Name,
		Location:    newSchool.Location,
		CreatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	}

	result, err := c.Repo.CreateSchool(SchoolToInsert)
	if err != nil {
		return nil, err
	}

	c.Validator.ClearErrors()
	return result, nil
}

func (c *SchoolService) UpdateSchool(id string, updatedData model.SchoolInput) (*model.School, error) {
	c.Validator.Validate(updatedData.Name, []string{"IsString", "Length:<25"})
	c.Validator.Validate(*updatedData.Location, []string{"IsString", "Length:<50"})

	validationErrors := c.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return nil, errors.New(errorMessage)
	}

	existingSchool, err := c.Repo.GetSchoolByID(id)
	if err != nil {
		return nil, err
	}

	timestamp := time.Now().String()
	newSchool := model.School{
		ID:          existingSchool.ID,
		Name:        updatedData.Name,
		Location:    updatedData.Location,
		CreatedAt:   existingSchool.CreatedAt,
		UpdatedAt:   &timestamp,
		SoftDeleted: existingSchool.SoftDeleted,
	}

	result, err := c.Repo.UpdateSchool(id, newSchool)
	if err != nil {
		return nil, err
	}

	c.Validator.ClearErrors()
	return result, nil
}

func (c *SchoolService) DeleteSchool(id string) error {
	c.Validator.Validate(id, []string{"IsUUID"})

	validationErrors := c.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return errors.New(errorMessage)
	}

	err := c.Repo.DeleteSchoolByID(id)
	if err != nil {
		return err
	}

	c.Validator.ClearErrors()
	return nil
}

func (c *SchoolService) GetSchoolById(id string) (*model.School, error) {
	c.Validator.Validate(id, []string{"IsUUID"})

	validationErrors := c.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		return nil, errors.New(errorMessage)
	}

	School, err := c.Repo.GetSchoolByID(id)
	if err != nil {
		return nil, err
	}

	c.Validator.ClearErrors()
	return School, nil
}

func (c *SchoolService) ListSchools() ([]*model.School, error) {
	Schools, err := c.Repo.ListSchools()
	if err != nil {
		return nil, err
	}

	return Schools, nil
}
