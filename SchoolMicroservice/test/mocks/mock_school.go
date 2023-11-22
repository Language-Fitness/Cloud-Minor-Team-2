package mocks

import (
	"example/graph/model"
	"time"
)

var Location = "This is a sample location."
var UpdatedLocation = "This is an updated sample location."
var Timestamp = time.Now().String()
var SoftDeleted = false

var MockCreateInput = model.SchoolInput{
	Name:     "Sample School",
	Location: &Location,
}

var MockUpdateInput = model.SchoolInput{
	Name:     "Sample School",
	Location: &UpdatedLocation,
}

var MockSchool = model.School{
	ID:          "3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
	Name:        "Sample School",
	Location:    &Location,
	CreatedAt:   &Timestamp,
	SoftDeleted: &SoftDeleted,
}

var MockUpdatedSchool = model.School{
	ID:          "3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
	Name:        "Sample School",
	Location:    &UpdatedLocation,
	CreatedAt:   &Timestamp,
	UpdatedAt:   &Timestamp,
	SoftDeleted: &SoftDeleted,
}
