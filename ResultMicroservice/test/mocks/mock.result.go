package mocks

import "ResultMicroservice/graph/model"

var ResultExerciseID = "1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n5o6p"
var ResultUserID = "a1b2c3d4-e5f6-g7h8-i9j0k1l2m3n4"
var ResultInput = "Sample user input"
var ResultResult = true
var ResultCreatedAt = "2024-01-04T09:45:00Z"
var ResultUpdatedAt = "2024-01-04T10:30:00Z"
var ResultSoftDeleted = false
var ClassID = "5a7bd776-6373-4e29-8aba-5b7571bdb7f5"
var ModuleID = "2a7bd276-2373-4e22-2aba-2b7521bdb2f6"

var MockResultInput = model.InputResult{
	ExerciseID: ResultExerciseID,
	UserID:     ResultUserID,
	ClassID:    ClassID,
	ModuleID:   ModuleID,
	Input:      ResultInput,
	Result:     ResultResult,
}

var MockResult = model.Result{
	ID:          "1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n5o6p",
	ExerciseID:  ResultExerciseID,
	UserID:      ResultUserID,
	ClassID:     ClassID,
	ModuleID:    ModuleID,
	Input:       ResultInput,
	Result:      ResultResult,
	CreatedAt:   ResultCreatedAt,
	UpdatedAt:   ResultUpdatedAt,
	SoftDeleted: ResultSoftDeleted,
}

var MockResultFilter = model.ResultFilter{
	SoftDelete: &ResultSoftDeleted,
	ExerciseID: &ResultExerciseID,
	UserID:     &ResultUserID,
	ClassID:    &ClassID,
	ModuleID:   &ModuleID,
	Input:      &ResultInput,
	Result:     &ResultResult,
}

var MockResultInfo = model.ResultInfo{
	ID:         "1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n5o6p",
	ExerciseID: ResultExerciseID,
	UserID:     ResultUserID,
	ClassID:    ClassID,
	ModuleID:   ModuleID,
	Input:      ResultInput,
	Result:     ResultResult,
}

var MockDeletedResult = model.Result{
	ID:          "1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n5o6p",
	ExerciseID:  ResultExerciseID,
	UserID:      ResultUserID,
	ClassID:     ClassID,
	ModuleID:    ModuleID,
	Input:       ResultInput,
	Result:      ResultResult,
	CreatedAt:   ResultCreatedAt,
	UpdatedAt:   ResultUpdatedAt,
	SoftDeleted: true,
}
