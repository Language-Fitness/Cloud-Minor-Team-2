package mocks

import "ExerciseMicroservice/graph/model"

var ExerciseName = "Sample Exercise"
var ExerciseQuestion = "This is a sample question."
var ExerciseAnswers = "Option A, Option B, Option C"
var ExercisePosCorrectAnswer = 1
var ExerciseDifficulty = model.LanguageLevelB1
var ExerciseCreatedAt = "2024-01-03T12:00:00Z"
var ExerciseUpdatedAt = "2024-01-03T13:30:00Z"
var ExerciseSoftDeleted = false
var ExerciseMadeBy = "5978e6ba-d199-426d-a643-3f7b3509b0d5"
var ExerciseID = "3a3bd756-6353-4e29-8aba-5b3531bdb9ef"
var ExerciseID2 = "3a3bd756-6353-4e29-8aba-5b3531bdb9f0"
var ClassID = "5a7bd776-6373-4e29-8aba-5b7571bdb7f5"
var ModuleID = "2a7bd276-2373-4e22-2aba-2b7521bdb2f6"

var MockExerciseInput = model.ExerciseInput{
	ClassID:          ClassID,
	ModuleID:         ModuleID,
	Name:             ExerciseName,
	Question:         ExerciseQuestion,
	Answers:          ExerciseAnswers,
	PosCorrectAnswer: ExercisePosCorrectAnswer,
	Difficulty:       ExerciseDifficulty,
}

var MockExercise = model.Exercise{
	ID:               ExerciseID,
	ClassID:          ClassID,
	ModuleID:         ModuleID,
	Name:             ExerciseName,
	Question:         ExerciseQuestion,
	Answers:          ExerciseAnswers,
	PosCorrectAnswer: ExercisePosCorrectAnswer,
	Difficulty:       ExerciseDifficulty,
	CreatedAt:        ExerciseCreatedAt,
	UpdatedAt:        ExerciseUpdatedAt,
	SoftDeleted:      ExerciseSoftDeleted,
	MadeBy:           ExerciseMadeBy,
}

var MockExerciseFilter = model.ExerciseFilter{
	Name:       &ExerciseName,
	SoftDelete: &ExerciseSoftDeleted,
	Difficulty: &ExerciseDifficulty,
	ClassID:    &MockExercise.ClassID,
	ModuleID:   &MockExercise.ModuleID,
	MadeBy:     &ExerciseMadeBy,
}

var MockDeletedExercise = model.Exercise{
	ID:               ExerciseID,
	ClassID:          ClassID,
	ModuleID:         ModuleID,
	Name:             ExerciseName,
	Question:         ExerciseQuestion,
	Answers:          ExerciseAnswers,
	PosCorrectAnswer: ExercisePosCorrectAnswer,
	Difficulty:       ExerciseDifficulty,
	CreatedAt:        ExerciseCreatedAt,
	UpdatedAt:        ExerciseUpdatedAt,
	SoftDeleted:      true,
	MadeBy:           ExerciseMadeBy,
}
