package service

import (
	"ResultMicroservice/graph/model"
	"ResultMicroservice/internal/auth"
	"ResultMicroservice/internal/database"
	"ResultMicroservice/internal/repository"
	"ResultMicroservice/internal/validation"
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

const ValidationPrefix = "Validation errors: "

// IResultService GOLANG INTERFACE
// Implements CRUD methods for queries and mutations on Result.
type IResultService interface {
	CreateResult(token string, newResult model.InputResult) (*model.Result, error)
	UpdateResult(token string, id string, updateData model.InputResult) (*model.Result, error)
	DeleteResult(token string, id string) error
	GetResultById(token string, id string) (*model.Result, error)
	ListResults(token string) ([]*model.Result, error)
	SoftDeleteByUser(token string, userID string) (string, bool, error)
	SoftDeleteByClass(token string, classID string) (string, bool, error)
	SoftDeleteByModule(token string, moduleID string) (string, bool, error)
	DeleteByUser(token string, userID string) (string, bool, error)
	DeleteByClass(token string, classID string) (string, bool, error)
	DeleteByModule(token string, moduleID string) (string, bool, error)
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

func (r *ResultService) ListResults(token string) ([]*model.Result, error) {
	err := r.ResultPolicy.ListResult(token)
	if err != nil {
		return nil, err
	}

	// Implement the logic to list results
	panic("implement me")
}

func (r *ResultService) CreateResult(token string, newResult model.InputResult) (*model.Result, error) {
	err := r.ResultPolicy.CreateResult(token)
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
		UserID:      newResult.UserID,
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
	r.Validator.Validate(id, []string{"IsUUID"})

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

func (r *ResultService) DeleteResult(token string, id string) error {
	err := r.ResultPolicy.DeleteResult(token, id)
	if err != nil {
		return err
	}

	r.Validator.Validate(id, []string{"IsUUID"})

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return errors.New(errorMessage)
	}

	err2 := r.Repo.DeleteResultByID(id)
	if err2 != nil {
		return err2
	}

	return nil
}

func (r *ResultService) GetResultById(token string, id string) (*model.Result, error) {
	err := r.ResultPolicy.GetResultByID(token, id)
	if err != nil {
		return nil, err
	}

	r.Validator.Validate(id, []string{"IsUUID"})

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	result, err := r.Repo.GetResultByID(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Saga grpc methods

func (r *ResultService) SoftDeleteByUser(token string, userID string) (string, bool, error) {
	err := r.ResultPolicy.SoftDeleteByUser(token, userID)
	if err != nil {
		return userID, false, err
	}

	r.Validator.Validate(userID, []string{"IsUUID"})

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return userID, false, errors.New(errorMessage)
	}

	err2 := r.Repo.SoftDeleteByUser(userID)
	if err2 != nil {
		return userID, false, err2
	}

	return userID, true, nil
}

func (r *ResultService) SoftDeleteByClass(token string, classID string) (string, bool, error) {
	err := r.ResultPolicy.SoftDeleteByClass(token, classID)
	if err != nil {
		return classID, false, err
	}

	r.Validator.Validate(classID, []string{"IsUUID"})

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return classID, false, errors.New(errorMessage)
	}

	err2 := r.Repo.SoftDeleteByClass(classID)
	if err2 != nil {
		return classID, false, err2
	}

	return classID, true, nil
}

func (r *ResultService) SoftDeleteByModule(token string, moduleID string) (string, bool, error) {
	err := r.ResultPolicy.SoftDeleteByModule(token, moduleID)
	if err != nil {
		return moduleID, false, err
	}

	r.Validator.Validate(moduleID, []string{"IsUUID"})

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return moduleID, false, errors.New(errorMessage)
	}

	err2 := r.Repo.SoftDeleteByModule(moduleID)
	if err2 != nil {
		return moduleID, false, err2
	}

	return moduleID, true, nil
}

func (r *ResultService) DeleteByUser(token string, userID string) (string, bool, error) {
	err := r.ResultPolicy.DeleteByUser(token, userID)
	if err != nil {
		return userID, false, err
	}

	r.Validator.Validate(userID, []string{"IsUUID"})

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return userID, false, errors.New(errorMessage)
	}

	err2 := r.Repo.DeleteByUser(userID)
	if err2 != nil {
		return userID, false, err2
	}

	return userID, true, nil
}

func (r *ResultService) DeleteByClass(token string, classID string) (string, bool, error) {
	err := r.ResultPolicy.DeleteByClass(token, classID)
	if err != nil {
		return classID, false, err
	}

	r.Validator.Validate(classID, []string{"IsUUID"})

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return classID, false, errors.New(errorMessage)
	}

	err2 := r.Repo.DeleteByClass(classID)
	if err2 != nil {
		return classID, false, err2
	}

	return classID, true, nil
}

func (r *ResultService) DeleteByModule(token string, moduleID string) (string, bool, error) {
	err := r.ResultPolicy.DeleteByModule(token, moduleID)
	if err != nil {
		return moduleID, false, err
	}

	r.Validator.Validate(moduleID, []string{"IsUUID"})

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return moduleID, false, errors.New(errorMessage)
	}

	err2 := r.Repo.DeleteByModule(moduleID)
	if err2 != nil {
		return moduleID, false, err2
	}

	return moduleID, true, nil
}

func (r *ResultService) ValidateResult(result *model.InputResult) {
	r.Validator.Validate(result.ExerciseID, []string{"IsUUID"})
	r.Validator.Validate(result.UserID, []string{"IsUUID"})
	r.Validator.Validate(result.ClassID, []string{"IsUUID"})
	r.Validator.Validate(result.ModuleID, []string{"IsUUID"})
	r.Validator.Validate(result.Input, []string{"IsString"})
	r.Validator.Validate(result.Result, []string{"IsString"})
}
