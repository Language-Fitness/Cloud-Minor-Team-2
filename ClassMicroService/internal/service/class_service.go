package service

import (
	"Class/graph/model"
	"Class/internal/auth"
	"Class/internal/database"
	"Class/internal/helper"
	"Class/internal/repository"
	"Class/internal/validation"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

const ValidationPrefix = "Validation errors: "

// IClassService GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on Class.
type IClassService interface {
	CreateClass(token string, newClass model.ClassInput) (*model.Class, error)
	UpdateClass(token string, id string, updatedData model.ClassInput) (*model.Class, error)
	DeleteClass(token string, id string, deleteFlag bool) error
	GetClassById(token string, id string) (*model.Class, error)
	ListClasses(token string, filter *model.ListClassFilter, paginate *model.Paginator) ([]*model.ClassInfo, error)
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

	validateNewClass(c.Validator, newClass)
	validationErrors := c.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
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

	validateUpdatedClass(c.Validator, id, updatedData)
	validationErrors := c.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
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

func (c *ClassService) DeleteClass(token string, id string, deleteFlag bool) error {
	existingClass, err := c.Policy.DeleteClass(token, id)
	if err != nil {
		return err
	}

	softDelete := deleteFlag
	existingClass.SoftDeleted = &softDelete

	err = c.Repo.DeleteClass(id, *existingClass)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClassService) GetClassById(token string, id string) (*model.Class, error) {
	existingClass, err := c.Policy.GetClass(token, id)
	if err != nil {
		return nil, err
	}

	return existingClass, nil
}

func (c *ClassService) ListClasses(token string, filter *model.ListClassFilter, paginate *model.Paginator) ([]*model.ClassInfo, error) {
	err := c.Policy.ListClasses(token)
	if err != nil {
		return nil, err
	}

	validateListClassFilter(c.Validator, filter, paginate)
	validationErrors := c.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		c.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	bsonFilter := buildBsonFilter(c.Policy, token, filter)

	paginateOptions := options.Find().
		SetSkip(int64(paginate.Step)).
		SetLimit(int64(paginate.Amount))

	classes, err := c.Repo.ListClasses(bsonFilter, paginateOptions)
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func validateListClassFilter(validator validation.IValidator, filter *model.ListClassFilter, paginate *model.Paginator) {
	validator.Validate(filter.SoftDelete, []string{"IsNull", "IsBoolean"}, "Filter softDelete")
	if helper.IsNil(filter.Name) == false {
		validator.Validate(helper.DereferenceArrayIfNeeded(
			filter.Name.Input),
			[]string{"IsNull", "ArrayType:string"},
			"Filter Name input")
	}
	validator.Validate(filter.ModuleID, []string{"IsNull", "IsUUID"}, "Filter ModuleID")
	validator.Validate(filter.MadeBy, []string{"IsNull", "IsUUID"}, "Filter Made_By")
	validator.Validate(paginate.Amount, []string{"IsInt", "Size:>0", "Size:<101"}, "Paginate Amount")
	validator.Validate(paginate.Step, []string{"IsInt", "Size:>=0"}, "Paginate Step")
}

func validateUpdatedClass(validator validation.IValidator, id string, updatedData model.ClassInput) {
	validator.Validate(id, []string{"IsUUID"}, "ID")
	validator.Validate(updatedData.ModuleID, []string{"IsUUID"}, "Module ID")
	validator.Validate(updatedData.Name, []string{"IsString", "Length:<25"}, "Name")
	validator.Validate(updatedData.Description, []string{"IsString", "Length:<50"}, "Description")
}

func validateNewClass(validator validation.IValidator, newClass model.ClassInput) {
	validator.Validate(newClass.ModuleID, []string{"IsUUID"}, "Module ID")
	validator.Validate(newClass.Name, []string{"IsString", "Length:<25"}, "Name")
	validator.Validate(newClass.Description, []string{"IsString", "Length:<50"}, "Description")
}

func buildBsonFilter(policy auth.IPolicy, token string, filter *model.ListClassFilter) bson.D {
	bsonFilter := bson.D{}

	if policy.HasPermissions(token, "filter_class_softDelete") == true && !helper.IsNil(filter.SoftDelete) {
		bsonFilter = append(bsonFilter, bson.E{Key: "softdeleted", Value: helper.DereferenceIfNeeded(filter.SoftDelete)})
	} else {
		bsonFilter = append(bsonFilter, bson.E{Key: "softdeleted", Value: false})
	}

	if policy.HasPermissions(token, "filter_class_module_id") == true && !helper.IsNil(filter.ModuleID) {
		bsonFilter = append(bsonFilter, bson.E{Key: "moduleid", Value: helper.DereferenceIfNeeded(filter.ModuleID)})
	}

	if policy.HasPermissions(token, "filter_class_name") == true && helper.IsNil(filter.Name) == false {
		bsonFilter = helper.AddFilter(
			bsonFilter, "name",
			string(filter.Name.Type),
			helper.DereferenceArrayIfNeeded(filter.Name.Input))
	}

	if policy.HasPermissions(token, "filter_class_difficulty") == true && !helper.IsNil(filter.Difficulty) {
		bsonFilter = append(bsonFilter, bson.E{Key: "difficulty", Value: helper.DereferenceIfNeeded(filter.Difficulty)})
	}

	if policy.HasPermissions(token, "filter_class_made_by") == true && !helper.IsNil(filter.MadeBy) {
		bsonFilter = append(bsonFilter, bson.E{Key: "madeby", Value: helper.DereferenceIfNeeded(filter.MadeBy)})
	}

	return bsonFilter
}
