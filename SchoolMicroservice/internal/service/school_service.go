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

// ISchoolService GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on School.
type ISchoolService interface {
	CreateSchool(token string, School model.SchoolInput) (*model.School, error)
	UpdateSchool(token string, id string, updatedData model.SchoolInput) (*model.School, error)
	DeleteSchool(token string, id string, filter *model.Filter) error
	GetSchoolById(token string, id string) (*model.School, error)
	ListSchools(token string) ([]*model.SchoolInfo, error)
}

// SchoolService GOLANG STRUCT
// Contains two interfaces for a Validator and a Repo.
type SchoolService struct {
	Validator validation.IValidator
	Repo      repository.ISchoolRepository
	Policy    auth.IPolicy
}

// NewSchoolService GOLANG FACTORY
// Returns a SchoolService implementing ISchoolService.
func NewSchoolService() ISchoolService {
	collection, _ := database.GetCollection()

	return &SchoolService{
		Validator: validation.NewValidator(),
		Repo:      repository.NewSchoolRepository(collection),
		Policy:    auth.NewPolicy(collection),
	}
}

func (s *SchoolService) CreateSchool(token string, newSchool model.SchoolInput) (*model.School, error) {
	sub, err := s.Policy.CreateSchool(token)
	if err != nil {
		return nil, err
	}

	s.Validator.Validate(newSchool.Name, []string{"IsString", "Length:<25"})
	s.Validator.Validate(newSchool.Location, []string{"IsString", "Length:<50"})
	s.Validator.Validate(newSchool.Location, []string{"IsString", "Length:<50"})

	validationErrors := s.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		s.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()
	softDeleted := false

	SchoolToInsert := &model.School{
		ID:          uuid.New().String(),
		Name:        newSchool.Name,
		Location:    newSchool.Location,
		MadeBy:      sub,
		CreatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	}

	result, err := s.Repo.CreateSchool(SchoolToInsert)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SchoolService) UpdateSchool(token string, id string, updatedData model.SchoolInput) (*model.School, error) {
	existingSchool, err := s.Policy.UpdateSchool(token, id)
	if err != nil {
		return nil, err
	}

	s.Validator.Validate(id, []string{"IsUUID"})
	s.Validator.Validate(updatedData.Name, []string{"IsString", "Length:<25"})
	s.Validator.Validate(updatedData.Location, []string{"IsString", "Length:<50"})

	validationErrors := s.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		s.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()
	newSchool := model.School{
		ID:          existingSchool.ID,
		Name:        updatedData.Name,
		Location:    updatedData.Location,
		MadeBy:      existingSchool.MadeBy,
		CreatedAt:   existingSchool.CreatedAt,
		UpdatedAt:   &timestamp,
		SoftDeleted: existingSchool.SoftDeleted,
	}

	result, err := s.Repo.UpdateSchool(id, newSchool)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SchoolService) DeleteSchool(token string, id string, filter *model.Filter) error {
	existingSchool, err := s.Policy.DeleteSchool(token, id)
	if err != nil {
		return err
	}

	if !*existingSchool.SoftDeleted {
		softDelete := true
		existingSchool.SoftDeleted = &softDelete

		err := s.Repo.SoftDeleteSchoolByID(id, *existingSchool)
		if err != nil {
			return err
		}
		return nil
	}

	if filter != nil && !*filter.SoftDelete {
		err := s.Repo.HardDeleteSchoolByID(id)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("school could not be deleted")
}

func (s *SchoolService) GetSchoolById(token string, id string) (*model.School, error) {
	existingSchool, err := s.Policy.GetSchool(token, id)
	if err != nil {
		return nil, err
	}

	return existingSchool, nil
}

func (s *SchoolService) ListSchools(token string) ([]*model.SchoolInfo, error) {
	err := s.Policy.ListSchools(token)
	if err != nil {
		return nil, err
	}

	Schools, err := s.Repo.ListSchools()
	if err != nil {
		return nil, err
	}

	return Schools, nil
}
