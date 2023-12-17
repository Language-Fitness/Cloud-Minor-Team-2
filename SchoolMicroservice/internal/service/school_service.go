package service

import (
	"errors"
	"example/graph/model"
	"example/internal/auth"
	"example/internal/database"
	"example/internal/helper"
	"example/internal/repository"
	"example/internal/validation"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

// ISchoolService GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on School.
type ISchoolService interface {
	CreateSchool(token string, School model.SchoolInput) (*model.School, error)
	UpdateSchool(token string, id string, updatedData model.SchoolInput) (*model.School, error)
	DeleteSchool(token string, id string, filter *model.ListSchoolFilter) error
	GetSchoolById(token string, id string) (*model.School, error)
	ListSchools(token string, filter *model.ListSchoolFilter, paginate *model.Paginator) ([]*model.SchoolInfo, error)
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

	s.Validator.Validate(newSchool.Name, []string{"IsString", "Length:<25"}, "Name")
	s.Validator.Validate(newSchool.Location, []string{"IsString", "Length:<50"}, "Location")

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

	s.Validator.Validate(id, []string{"IsUUID"}, "ID")
	s.Validator.Validate(updatedData.Name, []string{"IsString", "Length:<25"}, "Name")
	s.Validator.Validate(updatedData.Location, []string{"IsString", "Length:<50"}, "Location")

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

func (s *SchoolService) DeleteSchool(token string, id string, filter *model.ListSchoolFilter) error {
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

func (s *SchoolService) ListSchools(token string, filter *model.ListSchoolFilter, paginate *model.Paginator) ([]*model.SchoolInfo, error) {
	err := s.Policy.ListSchools(token)
	if err != nil {
		return nil, err
	}

	s.Validator.Validate(filter.SoftDelete, []string{"IsNull", "IsBoolean"}, "Filter softDelete")

	if !helper.IsNil(filter.Name) {
		s.Validator.Validate(helper.DereferenceArrayIfNeeded(
			filter.Name.Input),
			[]string{"IsNull", "ArrayType:string"},
			"Filter Name input")
	}

	if !helper.IsNil(filter.Location) {
		s.Validator.Validate(helper.DereferenceArrayIfNeeded(
			filter.Location.Input),
			[]string{"IsNull", "ArrayType:string"},
			"Filter Location input")
	}

	s.Validator.Validate(filter.MadeBy, []string{"IsNull", "IsUUID"}, "Filter Made_By")
	s.Validator.Validate(paginate.Amount, []string{"IsInt", "Size:>0", "Size:<101"}, "Paginate Amount")
	s.Validator.Validate(paginate.Step, []string{"IsInt", "Size:>=0"}, "Paginate Step")

	validationErrors := s.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		s.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	bsonFilter := bson.D{}
	if s.Policy.HasPermissions(token, "filter_school_softDelete") == true && !helper.IsNil(filter.SoftDelete) {
		bsonFilter = append(bsonFilter, bson.E{Key: "softdeleted", Value: helper.DereferenceIfNeeded(filter.SoftDelete)})
	}

	if s.Policy.HasPermissions(token, "filter_school_name") == true && !helper.IsNil(filter.Name) {
		bsonFilter = helper.AddFilter(
			bsonFilter, "name",
			string(filter.Name.Type),
			helper.DereferenceArrayIfNeeded(filter.Name.Input))
	}

	if s.Policy.HasPermissions(token, "filter_school_location") == true && !helper.IsNil(filter.Location) {
		bsonFilter = helper.AddFilter(
			bsonFilter,
			"location",
			string(filter.Location.Type),
			helper.DereferenceArrayIfNeeded(filter.Location.Input))
	}

	if s.Policy.HasPermissions(token, "filter_school_made_by") == true && !helper.IsNil(filter.MadeBy) {
		bsonFilter = append(bsonFilter, bson.E{Key: "madeby", Value: helper.DereferenceIfNeeded(filter.MadeBy)})
	}

	fmt.Println(bsonFilter)

	paginateOptions := options.Find().
		SetSkip(int64(paginate.Step)).
		SetLimit(int64(paginate.Amount))

	schools, err := s.Repo.ListSchools(bsonFilter, paginateOptions)
	if err != nil {
		return nil, err
	}

	return schools, nil
}
