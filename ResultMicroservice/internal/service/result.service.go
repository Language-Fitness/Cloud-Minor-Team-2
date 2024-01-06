package service

import (
	"ResultMicroservice/graph/model"
	"ResultMicroservice/internal/auth"
	"ResultMicroservice/internal/database"
	"ResultMicroservice/internal/helper"
	"ResultMicroservice/internal/repository"
	"ResultMicroservice/internal/validation"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"strings"
	"time"
)

const ValidationPrefix = "Validation errors: "

// IResultService GOLANG INTERFACE
// Implements CRUD methods for queries and mutations on Result.
type IResultService interface {
	CreateResult(token string, newResult model.InputResult) (*model.Result, error)
	UpdateResult(token string, id string, updateData model.InputResult) (*model.Result, error)
	DeleteResult(token string, id string) (*model.Result, error)
	GetResultById(token string, id string) (*model.Result, error)
	ListResults(token string, filter *model.ResultFilter, paginate *model.Paginator) ([]*model.Result, error)
}

// ResultService GOLANG STRUCT
// Contains two interfaces for a Validator and a Repo.
type ResultService struct {
	Validator    validation.IValidator
	Repo         repository.IResultRepository
	ResultPolicy auth.IResultPolicy
}

// NewResultService GOLANG FACTORY
// Returns a ResultService implementing IResultService.
func NewResultService() IResultService {
	collection, _ := database.GetCollection()

	return &ResultService{
		Validator:    validation.NewValidator(),
		Repo:         repository.NewResultRepository(collection),
		ResultPolicy: auth.NewResultPolicy(collection),
	}
}

func (r *ResultService) CreateResult(token string, newResult model.InputResult) (*model.Result, error) {
	id, err := r.ResultPolicy.CreateResult(token)
	if err != nil {
		return nil, err
	}

	r.ValidateResult(&newResult)
	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()

	resultToInsert := &model.Result{
		ID:          uuid.New().String(),
		ExerciseID:  newResult.ExerciseID,
		UserID:      id,
		ClassID:     newResult.ClassID,
		ModuleID:    newResult.ModuleID,
		Input:       newResult.Input,
		Result:      newResult.Result,
		CreatedAt:   timestamp,
		SoftDeleted: false,
	}

	result, err := r.Repo.CreateResult(resultToInsert)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ResultService) UpdateResult(token string, id string, updateData model.InputResult) (*model.Result, error) {
	result, err := r.ResultPolicy.UpdateResult(token, id)
	if err != nil {
		return nil, err
	}

	r.ValidateResult(&updateData)
	r.Validator.Validate(id, []string{"IsUUID"}, "ID")

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()

	newResult := model.Result{
		ID:          result.ID,
		ExerciseID:  updateData.ExerciseID,
		UserID:      updateData.UserID,
		ClassID:     updateData.ClassID,
		ModuleID:    updateData.ModuleID,
		Input:       updateData.Input,
		Result:      updateData.Result,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   timestamp,
		SoftDeleted: result.SoftDeleted,
	}

	updatedResult, err := r.Repo.UpdateResult(id, newResult)
	if err != nil {
		return nil, err
	}

	return updatedResult, nil
}

func (r *ResultService) DeleteResult(token string, id string) (*model.Result, error) {
	r.Validator.Validate(id, []string{"IsUUID"}, "ID")

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	//validate first because policy does not validate, and does a database request
	existingResult, err := r.ResultPolicy.DeleteResult(token, id)
	if err != nil {
		return nil, err
	}

	if existingResult.SoftDeleted {
		return nil, errors.New("result already soft deleted")
	}

	existingResult.SoftDeleted = true
	existingResult.UpdatedAt = time.Now().String()

	updatedResult, err := r.Repo.UpdateResult(id, *existingResult)
	if err != nil {
		return nil, err
	}

	return updatedResult, nil
}

func (r *ResultService) GetResultById(token string, id string) (*model.Result, error) {
	r.Validator.Validate(id, []string{"IsUUID"}, "ID")

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	ExistingResult, err := r.ResultPolicy.GetResultByID(token, id)
	if err != nil {
		return nil, err
	}

	return ExistingResult, nil
}

func (r *ResultService) ListResults(token string, filter *model.ResultFilter, paginate *model.Paginator) ([]*model.Result, error) {
	_, err := r.ResultPolicy.ListResult(token)
	if err != nil {
		return nil, err
	}

	validateListResultFilter(r.Validator, filter, paginate)
	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	bsonFilter, errs := buildBsonFilterForListResult(r.ResultPolicy, token, filter)
	if len(errs) > 0 {
		return nil, helper.AggregateErrors(errs)
	}

	paginateOptions := options.Find().
		SetSkip(int64(paginate.Step)).
		SetLimit(int64(paginate.Amount))

	results, err2 := r.Repo.ListResults(bsonFilter, paginateOptions)
	if err2 != nil {
		return nil, err2
	}

	return results, nil
}

func validateListResultFilter(validator validation.IValidator, filter *model.ResultFilter, paginate *model.Paginator) {
	validator.Validate(filter.SoftDelete, []string{"IsNull", "IsBoolean"}, "Filter SoftDelete")
	validator.Validate(filter.ExerciseID, []string{"IsNull", "IsString"}, "Filter ExerciseID")
	validator.Validate(filter.UserID, []string{"IsNull", "IsString"}, "Filter UserID")
	validator.Validate(filter.ClassID, []string{"IsNull", "IsString"}, "Filter ClassID")
	validator.Validate(filter.ModuleID, []string{"IsNull", "IsString"}, "Filter ModuleID")
	validator.Validate(paginate.Amount, []string{"IsInt", "Size:>0", "Size:<101"}, "Paginate Amount")
	validator.Validate(paginate.Step, []string{"IsInt", "Size:>=0"}, "Paginate Step")
}

func buildBsonFilterForListResult(policy auth.IResultPolicy, token string, filter *model.ResultFilter) (bson.D, []error) {
	bsonFilter := bson.D{}
	var errs []error

	appendCondition := func(key string, value interface{}, dbKey string) bool {
		if value != nil && !reflect.ValueOf(value).IsZero() && policy.HasPermissions(token, "filter_result_"+key) {
			bsonFilter = append(bsonFilter, bson.E{Key: dbKey, Value: value})
			return true
		} else if value != nil && !reflect.ValueOf(value).IsZero() && !policy.HasPermissions(token, "filter_result_"+key) {
			errs = append(errs, errors.New("invalid permissions for filter_result_"+key+" action, "))
			return false
		}
		return false
	}

	b := appendCondition("softdeleted", filter.SoftDelete, "softdeleted")
	if b == false {
		bsonFilter = append(bsonFilter, bson.E{Key: "softdeleted", Value: false})
	}
	appendCondition("exercise_id", filter.ExerciseID, "exerciseid")
	appendCondition("user_id", filter.UserID, "userid")
	appendCondition("class_id", filter.ClassID, "classid")
	appendCondition("module_id", filter.ModuleID, "moduleid")

	return bsonFilter, errs
}

func (r *ResultService) ValidateResult(result *model.InputResult) {
	r.Validator.Validate(result.ExerciseID, []string{"IsUUID"}, "ExerciseID")
	r.Validator.Validate(result.UserID, []string{"IsUUID"}, "UserID")
	r.Validator.Validate(result.ClassID, []string{"IsUUID"}, "ClassID")
	r.Validator.Validate(result.ModuleID, []string{"IsUUID"}, "ModuleID")
	r.Validator.Validate(result.Input, []string{"IsString"}, "Input")
	r.Validator.Validate(result.Result, []string{"IsBoolean"}, "Result")
}
