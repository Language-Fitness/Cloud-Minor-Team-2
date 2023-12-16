package service

import (
	"ExerciseMicroservice/graph/model"
	"ExerciseMicroservice/internal/auth"
	"ExerciseMicroservice/internal/repository"
	"ExerciseMicroservice/internal/validation"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

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
func NewExerciseService(collection *mongo.Collection) IExerciseService {
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

	e.Validator.Validate(newExercise.ClassID, []string{"IsString"}, "ClassID")
	e.Validator.Validate(newExercise.Name, []string{"IsString", "Length:<25"}, "Name")
	e.Validator.Validate(newExercise.Question, []string{"IsString", "Length:<50"}, "Question")
	e.Validator.Validate(newExercise.Answers, []string{"IsArray"}, "Answers")
	e.Validator.Validate(newExercise.PosCorrectAnswer, []string{"IsInt"}, "PosCorrectAnswer")
	e.Validator.Validate(newExercise.QuestionTypeID, []string{"IsString"}, "QuestionTypeID")
	e.Validator.Validate(newExercise.Difficulty, []string{"IsInt"}, "Difficulty")
	e.Validator.Validate(newExercise.ModuleID, []string{"IsString"}, "ModuleID")

	validationErrors := e.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
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

	e.Validator.ClearErrors()
	return result, nil
}

func (e *ExerciseService) UpdateExercise(token string, id string, updateData model.ExerciseInput) (*model.Exercise, error) {
	existingExercise, err := e.Policy.UpdateExercise(token, id)
	if err != nil {
		return nil, err
	}

	// Validate exercise update input fields here...
	e.Validator.Validate(updateData.ClassID, []string{"IsString"}, "ClassID")
	e.Validator.Validate(updateData.Name, []string{"IsString", "Length:<25"}, "Name")
	e.Validator.Validate(updateData.Question, []string{"IsString", "Length:<50"}, "Question")
	e.Validator.Validate(updateData.Answers, []string{"IsArray"}, "Answers")
	e.Validator.Validate(updateData.PosCorrectAnswer, []string{"IsInt"}, "PosCorrectAnswer")
	e.Validator.Validate(updateData.QuestionTypeID, []string{"IsString"}, "QuestionTypeID")
	e.Validator.Validate(updateData.Difficulty, []string{"IsInt"}, "Difficulty")
	e.Validator.Validate(updateData.ModuleID, []string{"IsString"}, "ModuleID")

	validationErrors := e.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
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

	e.Validator.ClearErrors()
	return result, nil
}

func (e *ExerciseService) DeleteExercise(token string, id string, filter *model.ExerciseFilter) error {
	isAdmin, existingExercise, err := e.Policy.DeleteExercise(token, id)
	if err != nil {
		return err
	}

	if !existingExercise.SoftDeleted {
		existingExercise.SoftDeleted = true

		err := e.Repo.SoftDeleteExerciseByID(id, *existingExercise)
		if err != nil {
			return err
		}
		return nil
	}

	if isAdmin && filter != nil && !*filter.SoftDelete {
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
	// Validate filter input fields here...
	e.Validator.Validate(filter.SoftDelete, []string{"IsNull", "IsBoolean"}, "Filter SoftDelete")
	e.Validator.Validate(filter.Name, []string{"IsNull", "IsString"}, "Filter Name")
	e.Validator.Validate(filter.Difficulty, []string{"IsNull", "IsFloat64"}, "Filter Difficulty")
	e.Validator.Validate(filter.QuestionTypeID, []string{"IsNull", "IsString"}, "Filter QuestionTypeID")
	e.Validator.Validate(filter.ClassID, []string{"IsNull", "IsString"}, "Filter ClassID")
	e.Validator.Validate(filter.ModuleID, []string{"IsNull", "IsString"}, "Filter ModuleID")
	e.Validator.Validate(filter.MadeBy, []string{"IsNull", "IsString"}, "Filter MadeBy")

	//todo reduce cognitive complexity to many if statements (sonarlint)
	validationErrors := e.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		e.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	// Check if the user has permission to list exercises
	if !e.Policy.HasPermissions(token, "list_exercises") {
		return nil, errors.New("unauthorized: user does not have permission to list exercises")
	}

	bsonFilter := bson.D{}

	// Add conditions to bsonFilter based on the filter fields...
	if e.Policy.HasPermissions(token, "filter_exercise_SoftDelete") && filter.SoftDelete != nil {
		bsonFilter = append(bsonFilter, bson.E{Key: "softdeleted", Value: *filter.SoftDelete})
	}

	if e.Policy.HasPermissions(token, "filter_exercise_Name") && filter.Name != nil {
		bsonFilter = append(bsonFilter, bson.E{Key: "name", Value: *filter.Name})
	}

	if e.Policy.HasPermissions(token, "filter_exercise_Difficulty") && filter.Difficulty != nil {
		bsonFilter = append(bsonFilter, bson.E{Key: "difficulty", Value: *filter.Difficulty})
	}

	if e.Policy.HasPermissions(token, "filter_exercise_QuestionTypeID") && filter.QuestionTypeID != nil {
		bsonFilter = append(bsonFilter, bson.E{Key: "questiontypeid", Value: *filter.QuestionTypeID})
	}

	if e.Policy.HasPermissions(token, "filter_exercise_ClassID") && filter.ClassID != nil {
		bsonFilter = append(bsonFilter, bson.E{Key: "classid", Value: *filter.ClassID})
	}

	if e.Policy.HasPermissions(token, "filter_exercise_ModuleID") && filter.ModuleID != nil {
		bsonFilter = append(bsonFilter, bson.E{Key: "moduleid", Value: *filter.ModuleID})
	}

	if e.Policy.HasPermissions(token, "filter_exercise_MadeBy") && filter.MadeBy != nil {
		bsonFilter = append(bsonFilter, bson.E{Key: "madeby", Value: *filter.MadeBy})
	}

	paginateOptions := options.Find().
		SetSkip(int64(paginate.Step)).
		SetLimit(int64(paginate.Amount))

	exercises, err := e.Repo.ListExercises(bsonFilter, paginateOptions)
	if err != nil {
		return nil, err
	}

	return exercises, nil
}
