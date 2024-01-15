package service

import (
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"school/graph/model"
	"school/internal/auth"
	"school/internal/database"
	"school/internal/helper"
	"school/internal/repository"
	"school/internal/validation"
	"strings"
	"time"
)

const ValidationPrefix = "Validation errors: "

// ISchoolService GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on School.
type ISchoolService interface {
	CreateSchool(token string, School model.SchoolInput) (*model.School, error)
	UpdateSchool(token string, id string, updatedData model.SchoolInput) (*model.School, error)
	DeleteSchool(token string, id string, deleteFlag bool) error
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

	validateNewSchool(s.Validator, newSchool)
	validationErrors := s.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		s.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	if !helper.IsNil(newSchool.OpenaiKey) {
		err := helper.ValidateOpenAiKey(*newSchool.OpenaiKey)

		if err != nil {
			return nil, err
		}
	}

	timestamp := time.Now().String()
	softDeleted := false

	SchoolToInsert := &model.School{
		ID:              uuid.New().String(),
		Name:            newSchool.Name,
		Location:        newSchool.Location,
		MadeBy:          sub,
		JoinCode:        uuid.New().String(),
		HasOpenaiAccess: newSchool.HasOpenaiAccess,
		CreatedAt:       &timestamp,
		SoftDeleted:     &softDeleted,
	}

	if newSchool.HasOpenaiAccess && !helper.IsNil(newSchool.OpenaiKey) {
		SchoolToInsert.OpenaiKey = newSchool.OpenaiKey
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

	validateUpdatedSchool(s.Validator, id, updatedData)
	validationErrors := s.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		s.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	err = validateOpenAiAttr(&updatedData, existingSchool.OpenaiKey)
	if err != nil {
		return nil, err
	}

	timestamp := time.Now().String()
	newSchool := model.School{
		ID:              existingSchool.ID,
		Name:            updatedData.Name,
		Location:        updatedData.Location,
		HasOpenaiAccess: updatedData.HasOpenaiAccess,
		OpenaiKey:       updatedData.OpenaiKey,
		JoinCode:        existingSchool.JoinCode,
		MadeBy:          existingSchool.MadeBy,
		CreatedAt:       existingSchool.CreatedAt,
		UpdatedAt:       &timestamp,
		SoftDeleted:     existingSchool.SoftDeleted,
	}

	if updatedData.HasOpenaiAccess && !helper.IsNil(updatedData.OpenaiKey) {
		newSchool.OpenaiKey = updatedData.OpenaiKey
	}

	result, err := s.Repo.UpdateSchool(id, newSchool)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SchoolService) DeleteSchool(token string, id string, deleteFlag bool) error {
	existingSchool, err := s.Policy.DeleteSchool(token, id)
	if err != nil {
		return err
	}

	if !*existingSchool.SoftDeleted {
		softDelete := deleteFlag
		existingSchool.SoftDeleted = &softDelete

		err := s.Repo.DeleteSchool(id, *existingSchool)
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

	validateListSchoolFilter(s.Validator, filter, paginate)
	validationErrors := s.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		s.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	bsonFilter := buildBsonFilter(s.Policy, token, filter)

	paginateOptions := options.Find().
		SetSkip(int64(paginate.Step)).
		SetLimit(int64(paginate.Amount))

	schools, err := s.Repo.ListSchools(bsonFilter, paginateOptions)
	if err != nil {
		return nil, err
	}

	return schools, nil
}

func validateUpdatedSchool(validator validation.IValidator, id string, updatedData model.SchoolInput) {
	validator.Validate(id, []string{"IsUUID"}, "ID")
	validator.Validate(updatedData.Name, []string{"IsString", "Length:<25"}, "Name")
	validator.Validate(updatedData.Location, []string{"IsString", "Length:<50"}, "Location")
	validator.Validate(updatedData.HasOpenaiAccess, []string{"IsBoolean"}, "Has OpenAI Access")

	if !helper.IsNil(updatedData.OpenaiKey) {
		validator.Validate(*updatedData.OpenaiKey, []string{"IsString"}, "Open AI Key")
	}
}

func validateNewSchool(validator validation.IValidator, newSchool model.SchoolInput) {
	validator.Validate(newSchool.Name, []string{"IsString", "Length:<25"}, "Name")
	validator.Validate(newSchool.Location, []string{"IsString", "Length:<50"}, "Location")
	validator.Validate(newSchool.HasOpenaiAccess, []string{"IsBoolean"}, "Filter Has OpenAI Access")

	if newSchool.HasOpenaiAccess && !helper.IsNil(newSchool.OpenaiKey) {
		validator.Validate(*newSchool.OpenaiKey, []string{"IsString"}, "Open AI Key")
	}
}

func validateListSchoolFilter(validator validation.IValidator, filter *model.ListSchoolFilter, paginate *model.Paginator) {
	validator.Validate(filter.SoftDelete, []string{"IsNull", "IsBoolean"}, "Filter softDelete")
	validator.Validate(filter.MadeBy, []string{"IsNull", "IsUUID"}, "Filter Made_By")
	validator.Validate(filter.HasOpenaiAccess, []string{"IsNull", "IsBoolean"}, "Filter Has OpenAI Access")
	validator.Validate(paginate.Amount, []string{"IsInt", "Size:>0", "Size:<101"}, "Paginate Amount")
	validator.Validate(paginate.Step, []string{"IsInt", "Size:>=0"}, "Paginate Step")

	if !helper.IsNil(filter.Name) {
		validator.Validate(helper.DereferenceArrayIfNeeded(
			filter.Name.Input),
			[]string{"IsNull", "ArrayType:string"},
			"Filter Name input")
	}

	if !helper.IsNil(filter.Location) {
		validator.Validate(helper.DereferenceArrayIfNeeded(
			filter.Location.Input),
			[]string{"IsNull", "ArrayType:string"},
			"Filter Location input")
	}
}

func buildBsonFilter(policy auth.IPolicy, token string, filter *model.ListSchoolFilter) bson.D {
	bsonFilter := bson.D{}
	if policy.HasPermissions(token, "filter_school_softDelete") == true && !helper.IsNil(filter.SoftDelete) {
		bsonFilter = append(bsonFilter, bson.E{Key: "softdeleted", Value: helper.DereferenceIfNeeded(filter.SoftDelete)})
	}

	if policy.HasPermissions(token, "filter_school_made_by") == true && !helper.IsNil(filter.MadeBy) {
		bsonFilter = append(bsonFilter, bson.E{Key: "madeby", Value: helper.DereferenceIfNeeded(filter.MadeBy)})
	}

	if policy.HasPermissions(token, "filter_school_has_openai_access") == true && !helper.IsNil(filter.HasOpenaiAccess) {
		bsonFilter = append(bsonFilter, bson.E{Key: "hasopenaiaccess", Value: helper.DereferenceIfNeeded(filter.HasOpenaiAccess)})
	}

	if policy.HasPermissions(token, "filter_school_name") == true && !helper.IsNil(filter.Name) {
		bsonFilter = helper.AddFilter(
			bsonFilter, "name",
			string(filter.Name.Type),
			helper.DereferenceArrayIfNeeded(filter.Name.Input))
	}

	if policy.HasPermissions(token, "filter_school_location") == true && !helper.IsNil(filter.Location) {
		bsonFilter = helper.AddFilter(
			bsonFilter,
			"location",
			string(filter.Location.Type),
			helper.DereferenceArrayIfNeeded(filter.Location.Input))
	}
	return bsonFilter
}

func validateOpenAiAttr(updatedData *model.SchoolInput, existingKey *string) error {
	if updatedData.HasOpenaiAccess && existingKey != updatedData.OpenaiKey {
		err := helper.ValidateOpenAiKey(*updatedData.OpenaiKey)
		if err != nil {
			return err
		}
	}

	if !updatedData.HasOpenaiAccess {
		emptyStr := ""
		updatedData.OpenaiKey = &emptyStr
	}
	return nil
}
