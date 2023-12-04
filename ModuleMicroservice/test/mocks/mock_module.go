package mocks

import (
	"Module/graph/model"
	"time"
)

var Description = "This is a sample module."
var UpdatedDescription = "Updated Description"
var Difficulty = 1
var Category = "Sample Category"
var MadeBy = "Sample User"
var Private = false
var Key = "sample-key"
var timestamp = time.Now().String()
var SoftDeleted = false

var MockCreateInput = model.ModuleInput{
	Name:        "Sample Module",
	Description: Description,
	Difficulty:  Difficulty,
	Category:    Category,
	Private:     Private,
	Key:         &Key,
}

var MockUpdateInput = model.ModuleInput{
	Name:        "Sample Module",
	Description: UpdatedDescription,
	Difficulty:  Difficulty,
	Category:    Category,
	Private:     Private,
	Key:         &Key,
}

var MockModule = model.Module{
	ID:          "3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
	Name:        "Sample Module",
	Description: Description,
	Difficulty:  Difficulty,
	Category:    Category,
	MadeBy:      MadeBy,
	Private:     Private,
	Key:         &Key,
	CreatedAt:   &timestamp,
	SoftDeleted: &SoftDeleted,
}

var MockModuleInfo = model.ModuleInfo{
	ID:          "3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
	Name:        "Sample Module",
	Description: Description,
	Difficulty:  Difficulty,
	Category:    Category,
	MadeBy:      MadeBy,
	Private:     Private,
}

var MockUpdatedModule = model.Module{
	ID:          "3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
	Name:        "Sample Module",
	Description: UpdatedDescription,
	Difficulty:  Difficulty,
	Category:    Category,
	MadeBy:      MadeBy,
	Private:     Private,
	Key:         &Key,
	CreatedAt:   &timestamp,
	SoftDeleted: &SoftDeleted,
}
