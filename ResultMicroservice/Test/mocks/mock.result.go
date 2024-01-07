// Package mocks provides mock data for testing purposes.
package mocks

import (
	"ResultMicroservice/graph/model"
	"time"
)

var MockInputResult = model.InputResult{
	ExerciseID: "sample_exercise_id",
	UserID:     "sample_user_id",
	ClassID:    "sample_class_id",
	ModuleID:   "sample_module_id",
	Input:      "sample_input",
	Result:     "sample_result",
}

var MockResult = model.Result{
	ID:          "sample_result_id",
	ExerciseID:  "sample_exercise_id",
	UserID:      "sample_user_id",
	ClassID:     "sample_class_id",
	ModuleID:    "sample_module_id",
	Input:       "sample_input",
	Result:      "sample_result",
	CreatedAt:   time.Now().Format(time.RFC3339),
	UpdatedAt:   time.Now().Format(time.RFC3339),
	SoftDeleted: false,
}

var SoftDeletedMockResult = model.Result{
	ID:          "sample_result_id",
	ExerciseID:  "sample_exercise_id",
	UserID:      "sample_user_id",
	ClassID:     "sample_class_id",
	ModuleID:    "sample_module_id",
	Input:       "sample_input",
	Result:      "sample_result",
	CreatedAt:   time.Now().Format(time.RFC3339),
	UpdatedAt:   time.Now().Format(time.RFC3339),
	SoftDeleted: true,
}
