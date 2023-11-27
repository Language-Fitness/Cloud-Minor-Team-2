package service

import (
	"ResultMicroservice/graph/model"
	"ResultMicroservice/internal/database"
	"ResultMicroservice/internal/repository"
	"ResultMicroservice/internal/validation"
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

// IResultService GOLANG INTERFACE
// Implements CRUD methods for queries and mutations on Result.
type IResultService interface {
	CreateResult(newResult model.InputResult) (*model.Result, error)
	UpdateResult(id string, updateData model.InputResult) (*model.Result, error)
	DeleteResult(id string) error
	GetResultById(id string) (*model.Result, error)
	GetResultByExerciseId(id string) (*model.Result, error)
	GetResultByClassId(id string) ([]*model.Result, error)
}

// ResultService GOLANG STRUCT
// Contains two interfaces for a Validator and a Repo.
type ResultService struct {
	Validator validation.IValidator
	Repo      repository.IResultRepository
}

// NewResultService GOLANG FACTORY
// Returns a ResultService implementing IResultService.
func NewResultService() IResultService {
	collection, _ := database.GetCollection()

	return &ResultService{
		Validator: validation.NewValidator(),
		Repo:      repository.NewResultRepository(collection),
	}
}

// TODO: add user id
func (r *ResultService) CreateResult(newResult model.InputResult) (*model.Result, error) {
	r.ValidateResult(&newResult)

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
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

	r.Validator.ClearErrors()
	return result, nil
}

// TODO: add user id
func (r *ResultService) UpdateResult(id string, updateData model.InputResult) (*model.Result, error) {
	r.ValidateResult(&updateData)

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	existingResult, err := r.Repo.GetResultByID(id)
	if err != nil {
		return nil, err
	}

	timestamp := time.Now().String()

	newResult := model.Result{
		ID:          existingResult.ID,
		ExerciseID:  updateData.ExerciseID,
		UserID:      updateData.UserID,
		ClassID:     updateData.ClassID,
		ModuleID:    updateData.ModuleID,
		Input:       updateData.Input,
		Result:      updateData.Result,
		CreatedAt:   existingResult.CreatedAt,
		UpdatedAt:   timestamp,
		SoftDeleted: existingResult.SoftDeleted,
	}

	result, err := r.Repo.UpdateResult(id, newResult)
	if err != nil {
		return nil, err
	}

	r.Validator.ClearErrors()
	return result, nil
}

// TODO: add user id
func (r *ResultService) DeleteResult(id string) error {
	r.Validator.Validate(id, []string{"IsUUID"})

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return errors.New(errorMessage)
	}

	err := r.Repo.DeleteResultByID(id)
	if err != nil {
		return err
	}

	r.Validator.ClearErrors()
	return nil
}

// TODO: add user id
func (r *ResultService) GetResultById(id string) (*model.Result, error) {
	r.Validator.Validate(id, []string{"IsUUID"})

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	result, err := r.Repo.GetResultByID(id)
	if err != nil {
		return nil, err
	}

	r.Validator.ClearErrors()
	return result, nil
}

// TODO: add user id
func (r *ResultService) GetResultByExerciseId(id string) (*model.Result, error) {
	r.Validator.Validate(id, []string{"IsUUID"})

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	result, err := r.Repo.GetResultByExerciseId(id)
	if err != nil {
		return nil, err
	}

	r.Validator.ClearErrors()
	return result, nil
}

// TODO: add user id
func (r *ResultService) GetResultByClassId(id string) ([]*model.Result, error) {
	r.Validator.Validate(id, []string{"IsUUID"})

	validationErrors := r.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		r.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	result, err := r.Repo.GetResultByClassId(id)
	if err != nil {
		return nil, err
	}

	r.Validator.ClearErrors()
	return result, nil
}

func (r *ResultService) ValidateResult(result *model.InputResult) {
	r.Validator.Validate(result.ExerciseID, []string{"IsUUID"})
	r.Validator.Validate(result.UserID, []string{"IsUUID"})
	r.Validator.Validate(result.ClassID, []string{"IsUUID"})
	r.Validator.Validate(result.ModuleID, []string{"IsUUID"})
	r.Validator.Validate(result.Input, []string{"IsString"})
	r.Validator.Validate(result.Result, []string{"IsString"})
}
