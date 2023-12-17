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

// IClassService GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on Class.
type IClassService interface {
	CreateClass(token string, newClass model.ClassInput) (*model.Class, error)
	UpdateClass(token string, id string, updatedData model.ClassInput) (*model.Class, error)
	DeleteClass(token string, id string, filter *model.ListClassFilter) error
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

	c.Validator.Validate(newClass.ModuleID, []string{"IsUUID"}, "Module ID")
	c.Validator.Validate(newClass.Name, []string{"IsString", "Length:<25"}, "Name")
	c.Validator.Validate(newClass.Description, []string{"IsString", "Length:<50"}, "Description")

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

	c.Validator.Validate(updatedData.ModuleID, []string{"IsUUID"}, "Module ID")
	c.Validator.Validate(updatedData.Name, []string{"IsString", "Length:<25"}, "Name")
	c.Validator.Validate(updatedData.Description, []string{"IsString", "Length:<50"}, "Description")

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

func (c *ClassService) DeleteClass(token string, id string, filter *model.ListClassFilter) error {
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

func (c *ClassService) ListClasses(token string, filter *model.ListClassFilter, paginate *model.Paginator) ([]*model.ClassInfo, error) {
	err := c.Policy.ListClasses(token)
	if err != nil {
		return nil, err
	}

	c.Validator.Validate(filter.SoftDelete, []string{"IsNull", "IsBoolean"}, "Filter softDelete")
	if helper.IsNil(filter.Name) == false {
		c.Validator.Validate(helper.DereferenceArrayIfNeeded(
			filter.Name.Input),
			[]string{"IsNull", "ArrayType:string"},
			"Filter Name input")
	}
	c.Validator.Validate(filter.ModuleID, []string{"IsNull", "IsUUID"}, "Filter ModuleID")
	c.Validator.Validate(filter.MadeBy, []string{"IsNull", "IsUUID"}, "Filter Made_By")
	c.Validator.Validate(paginate.Amount, []string{"IsInt", "Size:>0", "Size:<101"}, "Paginate Amount")
	c.Validator.Validate(paginate.Step, []string{"IsInt", "Size:>=0"}, "Paginate Step")

	validationErrors := c.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		c.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	bsonFilter := bson.D{}

	if c.Policy.HasPermissions(token, "filter_class_softDelete") == true && !helper.IsNil(filter.SoftDelete) {
		bsonFilter = append(bsonFilter, bson.E{Key: "softdeleted", Value: helper.DereferenceIfNeeded(filter.SoftDelete)})
	} else {
		bsonFilter = append(bsonFilter, bson.E{Key: "softdeleted", Value: false})
	}

	if c.Policy.HasPermissions(token, "filter_class_module_id") == true && !helper.IsNil(filter.ModuleID) {
		bsonFilter = append(bsonFilter, bson.E{Key: "moduleid", Value: helper.DereferenceIfNeeded(filter.ModuleID)})
	}

	if c.Policy.HasPermissions(token, "filter_class_name") == true && helper.IsNil(filter.Name) == false {
		bsonFilter = helper.AddFilter(
			bsonFilter, "name",
			string(filter.Name.Type),
			helper.DereferenceArrayIfNeeded(filter.Name.Input))
	}

	if c.Policy.HasPermissions(token, "filter_class_difficulty") == true && !helper.IsNil(filter.Difficulty) {
		bsonFilter = append(bsonFilter, bson.E{Key: "difficulty", Value: helper.DereferenceIfNeeded(filter.Difficulty)})
	}

	if c.Policy.HasPermissions(token, "filter_class_made_by") == true && !helper.IsNil(filter.MadeBy) {
		bsonFilter = append(bsonFilter, bson.E{Key: "madeby", Value: helper.DereferenceIfNeeded(filter.MadeBy)})
	}

	fmt.Println(bsonFilter)

	paginateOptions := options.Find().
		SetSkip(int64(paginate.Step)).
		SetLimit(int64(paginate.Amount))

	classes, err := c.Repo.ListClasses(bsonFilter, paginateOptions)
	if err != nil {
		return nil, err
	}

	return classes, nil
}
