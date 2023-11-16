package mocks

import (
	"example/graph/model"
	"time"
)

var Description = "This is a sample class."
var UpdatedDescription = "This is an updated sample class."
var Difficulty = 1
var Timestamp = time.Now().String()
var SoftDeleted = false

var MockCreateInput = model.ClassInput{
	ModuleID:    "module-id",
	Name:        "Sample Class",
	Description: &Description,
	Difficulty:  &Difficulty,
}

var MockUpdateInput = model.ClassInput{
	ModuleID:    "module-id",
	Name:        "Sample Class",
	Description: &UpdatedDescription,
	Difficulty:  &Difficulty,
}

var MockClass = model.Class{
	ID:          "3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
	ModuleID:    "module-id",
	Name:        "Sample Class",
	Description: &Description,
	Difficulty:  &Difficulty,
	CreatedAt:   &Timestamp,
	SoftDeleted: &SoftDeleted,
}

var MockUpdatedClass = model.Class{
	ID:          "3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
	ModuleID:    "module-id",
	Name:        "Sample Class",
	Description: &UpdatedDescription,
	Difficulty:  &Difficulty,
	CreatedAt:   &Timestamp,
	UpdatedAt:   &Timestamp,
	SoftDeleted: &SoftDeleted,
}
