package service

import (
	"ExerciseMicroservice/graph/model"
	"ExerciseMicroservice/internal/auth"
	"ExerciseMicroservice/internal/database"
	"ExerciseMicroservice/internal/helper"
	"ExerciseMicroservice/internal/repository"
	"ExerciseMicroservice/internal/validation"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"strings"
	"time"
)

const ValidationPrefix = "Validation errors: "

// IExerciseService GOLANG INTERFACE
// Implements five CRUD methods for queries and mutations on Exercise.
type IExerciseService interface {
	CreateExercise(token string, newExercise model.ExerciseInput) (*model.Exercise, error)
	UpdateExercise(token string, id string, updateData model.ExerciseInput) (*model.Exercise, error)
	DeleteExercise(token string, id string) error
	UnDeleteExercise(token string, id string) error
	GetExerciseById(token string, id string) (*model.Exercise, error)
	ListExercises(token string, filter *model.ExerciseFilter, paginate *model.Paginator) ([]*model.ExerciseInfo, error)
}

// ExerciseService GOLANG STRUCT
// Contains two interfaces for a Validator and a Repo.
type ExerciseService struct {
	Validator validation.IValidator
	Repo      repository.IExerciseRepository
	Policy    auth.IExercisePolicy
}

// NewExerciseService GOLANG FACTORY
// Returns an ExerciseService implementing IExerciseService.
func NewExerciseService() IExerciseService {
	collection, _ := database.GetCollection()
	return &ExerciseService{
		Validator: validation.NewValidator(),
		Repo:      repository.NewExerciseRepository(collection),
		Policy:    auth.NewExercisePolicy(collection),
	}
}

func (e *ExerciseService) CreateExercise(token string, newExercise model.ExerciseInput) (*model.Exercise, error) {
	id, err := e.Policy.CreateExercise(token)
	if err != nil {
		return nil, err
	}

	err2 := validateAnswers(e.Validator, newExercise.Answers)
	if err2 != nil {
		return nil, err2
	}

	validateNewExercise(e.Validator, newExercise)
	validationErrors := e.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		e.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	var answerArray []*model.Answer

	for _, input := range newExercise.Answers {
		object := &model.Answer{
			Value:   input.Value,
			Correct: input.Correct,
		}

		answerArray = append(answerArray, object)
	}

	timestamp := time.Now().String()

	exerciseToInsert := &model.Exercise{
		ID:          uuid.New().String(),
		ClassID:     newExercise.ClassID,
		ModuleID:    newExercise.ModuleID,
		Name:        newExercise.Name,
		Question:    newExercise.Question,
		Answers:     answerArray,
		Difficulty:  newExercise.Difficulty,
		CreatedAt:   timestamp,
		SoftDeleted: false,
		MadeBy:      id,
	}

	result, err := e.Repo.CreateExercise(exerciseToInsert)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (e *ExerciseService) UpdateExercise(token string, id string, updateData model.ExerciseInput) (*model.Exercise, error) {
	validateUpdatedExercise(e.Validator, id, updateData)
	err := validateAnswers(e.Validator, updateData.Answers)
	if err != nil {
		return nil, err
	}

	validationErrors := e.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		e.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	//validate first because policy does not validate, and does a database request
	existingExercise, err := e.Policy.UpdateExercise(token, id)
	if err != nil {
		return nil, err
	}

	var answerArray []*model.Answer

	for _, input := range updateData.Answers {
		object := &model.Answer{
			Value:   input.Value,
			Correct: input.Correct,
		}

		answerArray = append(answerArray, object)
	}

	timestamp := time.Now().String()
	newExercise := model.Exercise{
		ID:          existingExercise.ID,
		ClassID:     updateData.ClassID,
		Name:        updateData.Name,
		Question:    updateData.Question,
		Answers:     answerArray,
		Difficulty:  updateData.Difficulty,
		CreatedAt:   existingExercise.CreatedAt,
		UpdatedAt:   timestamp,
		SoftDeleted: existingExercise.SoftDeleted,
	}

	result, err := e.Repo.UpdateExercise(id, newExercise)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (e *ExerciseService) DeleteExercise(token string, id string) error {
	e.Validator.Validate(id, []string{"IsUUID"}, "ID")

	valErrors := e.Validator.GetErrors()
	if len(valErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(valErrors, ", ")
		e.Validator.ClearErrors()
		return errors.New(errorMessage)
	}
	//validate first because policy does not validate, and does a database request
	_, existingExercise, err := e.Policy.DeleteExercise(token, id)
	if err != nil {
		return err
	}

	if existingExercise.SoftDeleted {
		return errors.New("exercise is already deleted")
	}

	existingExercise.SoftDeleted = true
	existingExercise.UpdatedAt = time.Now().String()

	_, err = e.Repo.UpdateExercise(id, *existingExercise)
	if err != nil {
		return err
	}

	return nil
}

func (e *ExerciseService) UnDeleteExercise(token string, id string) error {
	e.Validator.Validate(id, []string{"IsUUID"}, "ID")

	valErrors := e.Validator.GetErrors()
	if len(valErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(valErrors, ", ")
		e.Validator.ClearErrors()
		return errors.New(errorMessage)
	}
	//validate first because policy does not validate, and does a database request
	_, existingExercise, err := e.Policy.DeleteExercise(token, id)
	if err != nil {
		return err
	}

	if !existingExercise.SoftDeleted {
		return errors.New("exercise is not soft deleted")
	}

	existingExercise.SoftDeleted = false
	existingExercise.UpdatedAt = time.Now().String()

	_, err = e.Repo.UpdateExercise(id, *existingExercise)
	if err != nil {
		return err
	}

	return nil
}

func (e *ExerciseService) GetExerciseById(token string, id string) (*model.Exercise, error) {
	e.Validator.Validate(id, []string{"IsUUID"}, "ID")

	validationErrors := e.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		e.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	existingExercise, err := e.Policy.GetExercise(token, id)
	if err != nil {
		return nil, err
	}

	return existingExercise, nil
}

func (e *ExerciseService) ListExercises(token string, filter *model.ExerciseFilter, paginate *model.Paginator) ([]*model.ExerciseInfo, error) {
	_, err := e.Policy.ListExercises(token)
	if err != nil {
		return nil, err
	}

	validateListExerciseFilter(e.Validator, filter, paginate)
	validationErrors := e.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		e.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	bsonFilter, errs := buildBsonFilterForListExercise(e.Policy, token, filter)
	if len(errs) > 0 {
		return nil, helper.AggregateErrors(errs)
	}

	paginateOptions := options.Find().
		SetSkip(int64(paginate.Step)).
		SetLimit(int64(paginate.Amount))

	exercises, err2 := e.Repo.ListExercises(bsonFilter, paginateOptions)
	if err2 != nil {
		return nil, err2
	}

	return exercises, nil
}

func validateListExerciseFilter(validator validation.IValidator, filter *model.ExerciseFilter, paginate *model.Paginator) {
	validator.Validate(filter.SoftDelete, []string{"IsNull", "IsBoolean"}, "Filter SoftDelete")
	validator.Validate(filter.Name, []string{"IsNull", "IsString"}, "Filter Name")
	validator.Validate(filter.ClassID, []string{"IsNull", "IsUUID"}, "Filter ClassID")
	validator.Validate(filter.ModuleID, []string{"IsNull", "IsUUID"}, "Filter ModuleID")
	validator.Validate(filter.MadeBy, []string{"IsNull", "IsString"}, "Filter MadeBy")
	validator.Validate(paginate.Amount, []string{"IsInt", "Size:>0", "Size:<101"}, "Paginate Amount")
	validator.Validate(paginate.Step, []string{"IsInt", "Size:>=0"}, "Paginate Step")
}

func validateUpdatedExercise(validator validation.IValidator, id string, updatedData model.ExerciseInput) {
	validator.Validate(id, []string{"IsUUID"}, "ID")
	validator.Validate(updatedData.ClassID, []string{"IsUUID"}, "ClassID")
	validator.Validate(updatedData.Name, []string{"IsString", "Length:<50"}, "Name")
	validator.Validate(updatedData.Question, []string{"IsString", "Length:<100"}, "Question")
	validator.Validate(updatedData.ModuleID, []string{"IsUUID"}, "ModuleID")
}

func validateAnswers(validator validation.IValidator, answers []*model.AnswerInput) error {
	if len(answers) < 2 {
		return errors.New("exercise must have at least two answers")
	}

	var correctCount int
	for _, answer := range answers {
		validator.Validate(answer.Value, []string{"IsString", "Length:<100"}, "Answer Value")
		validator.Validate(answer.Correct, []string{"IsBoolean"}, "Answer Correct")

		if answer.Correct {
			correctCount++
		}
	}

	// Check conditions
	if correctCount == 0 {
		return errors.New("at least one answer must be correct")
	} else if correctCount > 1 {
		return errors.New("only one answer can be correct")
	}
	return nil
}

func validateNewExercise(validator validation.IValidator, newExercise model.ExerciseInput) {
	validator.Validate(newExercise.ClassID, []string{"IsUUID"}, "ClassID")
	validator.Validate(newExercise.Name, []string{"IsString", "Length:<50"}, "Name")
	validator.Validate(newExercise.Question, []string{"IsString", "Length:<100"}, "Question")
	validator.Validate(newExercise.ModuleID, []string{"IsUUID"}, "ModuleID")
}

func buildBsonFilterForListExercise(policy auth.IExercisePolicy, token string, filter *model.ExerciseFilter) (bson.D, []error) {
	bsonFilter := bson.D{}
	//list of errors
	var errs []error

	fmt.Println(token)

	appendCondition := func(key string, value interface{}, dbKey string) bool {

		if value != nil && !reflect.ValueOf(value).IsZero() && policy.HasPermissions(token, "filter_exercise_"+key) {
			bsonFilter = append(bsonFilter, bson.E{Key: dbKey, Value: helper.DereferenceIfNeeded(value)})
			return true
		} else if value != nil && !reflect.ValueOf(value).IsZero() && !policy.HasPermissions(token, "filter_exercise_"+key) {
			errs = append(errs, errors.New("invalid permissions for filter_exercise_"+key+" action, "))
			return false
		}
		return false
	}

	b := appendCondition("softDelete", filter.SoftDelete, "softdeleted")
	if b == false {
		bsonFilter = append(bsonFilter, bson.E{Key: "softdeleted", Value: false})
	}
	appendCondition("name", filter.Name, "name")
	appendCondition("difficulty", filter.Difficulty, "difficulty")
	appendCondition("class_id", filter.ClassID, "classid")
	appendCondition("module_id", filter.ModuleID, "moduleid")
	appendCondition("made_by", filter.MadeBy, "madeby")

	return bsonFilter, errs
}
