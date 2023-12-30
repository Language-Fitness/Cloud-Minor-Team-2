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
	DeleteExercise(token string, id string, filter *model.ExerciseFilter) error
	GetExerciseById(token string, id string) (*model.Exercise, error)
	ListExercises(token string, filter *model.ExerciseFilter, paginate *model.Paginator) ([]*model.Exercise, error)
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
	sub, err := e.Policy.CreateExercise(token)
	if err != nil {
		return nil, err
	}

	validateNewExercise(e.Validator, newExercise)
	validationErrors := e.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		e.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()

	exerciseToInsert := &model.Exercise{
		ID:               uuid.New().String(),
		ClassID:          newExercise.ClassID,
		Name:             newExercise.Name,
		Question:         newExercise.Question,
		Answers:          newExercise.Answers,
		PosCorrectAnswer: newExercise.PosCorrectAnswer,
		QuestionTypeID:   newExercise.QuestionTypeID,
		Difficulty:       newExercise.Difficulty,
		CreatedAt:        timestamp,
		SoftDeleted:      false,
		MadeBy:           sub,
	}

	result, err := e.Repo.CreateExercise(exerciseToInsert)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (e *ExerciseService) UpdateExercise(token string, id string, updateData model.ExerciseInput) (*model.Exercise, error) {
	existingExercise, err := e.Policy.UpdateExercise(token, id)
	if err != nil {
		return nil, err
	}

	validateUpdatedExercise(e.Validator, id, updateData)
	validationErrors := e.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := ValidationPrefix + strings.Join(validationErrors, ", ")
		e.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()
	newExercise := model.Exercise{
		ID:               existingExercise.ID,
		ClassID:          updateData.ClassID,
		Name:             updateData.Name,
		Question:         updateData.Question,
		Answers:          updateData.Answers,
		PosCorrectAnswer: updateData.PosCorrectAnswer,
		QuestionTypeID:   updateData.QuestionTypeID,
		Difficulty:       updateData.Difficulty,
		CreatedAt:        existingExercise.CreatedAt,
		UpdatedAt:        timestamp,
		SoftDeleted:      existingExercise.SoftDeleted,
	}

	result, err := e.Repo.UpdateExercise(id, newExercise)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (e *ExerciseService) DeleteExercise(token string, id string, filter *model.ExerciseFilter) error {
	isAdmin, existingExercise, err := e.Policy.DeleteExercise(token, id)
	if err != nil {
		return err
	}

	if !existingExercise.SoftDeleted {
		existingExercise.SoftDeleted = true

		//TODO nog naar kijken  want delete moet update met hard delete variable zijn
		err := e.Repo.SoftDeleteExerciseByID(id, *existingExercise)
		if err != nil {
			return err
		}
		return nil
	}

	if isAdmin && filter != nil && !helper.IsNil(filter.SoftDelete) {
		err := e.Repo.HardDeleteExerciseByID(id)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("exercise could not be deleted")
}

func (e *ExerciseService) GetExerciseById(token string, id string) (*model.Exercise, error) {
	existingExercise, err := e.Policy.GetExercise(token, id)
	if err != nil {
		return nil, err
	}

	return existingExercise, nil
}

func (e *ExerciseService) ListExercises(token string, filter *model.ExerciseFilter, paginate *model.Paginator) ([]*model.Exercise, error) {
	_, err := e.Policy.ListExercises(token)
	if err != nil {
		return nil, err
	}

	fmt.Println("filter: ", filter)

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

	fmt.Println("bsonFilter: ", bsonFilter)

	exercises, err2 := e.Repo.ListExercises(bsonFilter, paginateOptions)
	if err2 != nil {
		return nil, err2
	}

	return exercises, nil
}

func validateListExerciseFilter(validator validation.IValidator, filter *model.ExerciseFilter, paginate *model.Paginator) {
	validator.Validate(filter.SoftDelete, []string{"IsNull", "IsBoolean"}, "Filter SoftDelete")
	if helper.IsNil(filter.Name) == false {
		validator.Validate(helper.DereferenceArrayIfNeeded(
			filter.Name.Input),
			[]string{"IsNull", "IsString"},
			"Filter Name input")
	}
	validator.Validate(filter.QuestionTypeID, []string{"IsNull", "IsString"}, "Filter QuestionTypeID")
	validator.Validate(filter.ClassID, []string{"IsNull", "IsString"}, "Filter ClassID")
	validator.Validate(filter.ModuleID, []string{"IsNull", "IsString"}, "Filter ModuleID")
	validator.Validate(filter.MadeBy, []string{"IsNull", "IsString"}, "Filter MadeBy")
	validator.Validate(paginate.Amount, []string{"IsInt", "Size:>0", "Size:<101"}, "Paginate Amount")
	validator.Validate(paginate.Step, []string{"IsInt", "Size:>=0"}, "Paginate Step")
}

func validateUpdatedExercise(validator validation.IValidator, id string, updatedData model.ExerciseInput) {
	validator.Validate(id, []string{"IsUUID"}, "ID")
	validator.Validate(updatedData.ClassID, []string{"IsString"}, "ClassID")
	validator.Validate(updatedData.Name, []string{"IsString", "Length:<25"}, "Name")
	validator.Validate(updatedData.Question, []string{"IsString", "Length:<50"}, "Question")
	validator.Validate(updatedData.Answers, []string{"IsString"}, "Answers")
	validator.Validate(updatedData.PosCorrectAnswer, []string{"IsInt"}, "PosCorrectAnswer")
	validator.Validate(updatedData.QuestionTypeID, []string{"IsString"}, "QuestionTypeID")
	validator.Validate(updatedData.ModuleID, []string{"IsString"}, "ModuleID")
}

// todo look if they need to be Dereferenced
func validateNewExercise(validator validation.IValidator, newExercise model.ExerciseInput) {
	validator.Validate(newExercise.ClassID, []string{"IsString"}, "ClassID")
	validator.Validate(newExercise.Name, []string{"IsString", "Length:<25"}, "Name")
	validator.Validate(newExercise.Question, []string{"IsString", "Length:<50"}, "Question")
	validator.Validate(newExercise.Answers, []string{"IsString"}, "Answers")
	validator.Validate(newExercise.PosCorrectAnswer, []string{"IsInt"}, "PosCorrectAnswer")
	validator.Validate(newExercise.QuestionTypeID, []string{"IsString"}, "QuestionTypeID")
	validator.Validate(newExercise.ModuleID, []string{"IsString"}, "ModuleID")
}

func buildBsonFilterForListExercise(policy auth.IExercisePolicy, token string, filter *model.ExerciseFilter) (bson.D, []error) {
	bsonFilter := bson.D{}
	//list of errors
	var errs []error

	appendCondition := func(key string, value interface{}) {

		if value != nil && !reflect.ValueOf(value).IsZero() && policy.HasPermissions(token, "filter_exercise_"+key) {
			bsonFilter = append(bsonFilter, bson.E{Key: key, Value: value})
		} else if value != nil && !reflect.ValueOf(value).IsZero() && !policy.HasPermissions(token, "filter_exercise_"+key) {
			errs = append(errs, errors.New("invalid permissions for filter_exercise_"+key+" action"))
		}
	}

	appendCondition("softdeleted", filter.SoftDelete)
	appendCondition("name", filter.Name)
	appendCondition("difficulty", filter.Difficulty)
	appendCondition("question_type_id", filter.QuestionTypeID)
	appendCondition("class_id", filter.ClassID)
	appendCondition("module_id", filter.ModuleID)
	appendCondition("madeby", filter.MadeBy)

	return bsonFilter, errs
}
