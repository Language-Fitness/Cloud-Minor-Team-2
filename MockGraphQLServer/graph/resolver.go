package graph

import "example/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	schools     []*model.School
	modules     []*model.Module
	classes     []*model.Class
	exercises   []*model.Exercise
	results     []*model.Result
	leaderboard []*model.LeaderboardRow
}

func (r *Resolver) init() {
	/// MODULES ///
	timestamp := "2023-01-01T00:00:00Z"
	softDeleted := false
	r.schools = append(r.schools, &model.School{
		ID:          "67bd70b8-8ac8-11ee-b9d1-0242ac120002",
		Name:        "Inholland",
		Location:    "Haarlem",
		MadeBy:      "bbf78bc0-942b-11ee-b9d1-0242ac120002",
		CreatedAt:   &timestamp,
		UpdatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	})

	/// MODULES ///
	r.modules = append(r.modules, &model.Module{
		ID:          "e41189cc-83cc-11ee-b962-0242ac120002",
		Name:        "Grammatica door Toon",
		Description: "Leer de basis van grammatica",
		Difficulty:  1,
		Category:    "Grammatica",
		MadeBy:      "3a6085b2-83cd-11ee-b962-0242ac120002",
		Private:     false,
		CreatedAt:   &timestamp,
		UpdatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	})

	key := "key"
	r.modules = append(r.modules, &model.Module{
		ID:          "25524ed4-83ce-11ee-b962-0242ac120002",
		Name:        "Werkwoordspelling door Toon",
		Description: "Leer de basis van grammatica",
		Difficulty:  1,
		Category:    "Werkwoordspelling",
		MadeBy:      "3a6085b2-83cd-11ee-b962-0242ac120002",
		Private:     false,
		Key:         &key,
		CreatedAt:   &timestamp,
		UpdatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	})

	/// CLASSES ///
	r.classes = append(r.classes, &model.Class{
		ID:          "f2f371e4-83cc-11ee-b962-0242ac120002",
		ModuleID:    "e41189cc-83cc-11ee-b962-0242ac120002",
		Name:        "Class A",
		Description: "Dit is Class A",
		Difficulty:  3,
		MadeBy:      "bbf78bc0-942b-11ee-b9d1-0242ac120002",
		CreatedAt:   &timestamp,
		UpdatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	})

	r.classes = append(r.classes, &model.Class{
		ID:          "7c1ee614-83ce-11ee-b962-0242ac120002",
		ModuleID:    "25524ed4-83ce-11ee-b962-0242ac120002",
		Name:        "Class A",
		Description: "Dit is Class A",
		Difficulty:  3,
		MadeBy:      "bbf78bc0-942b-11ee-b9d1-0242ac120002",
		CreatedAt:   &timestamp,
		UpdatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	})

	/// EXERCISES ///
	r.exercises = append(r.exercises, &model.Exercise{
		ID:               "af2cef6a-83ce-11ee-b962-0242ac120002",
		ClassID:          "f2f371e4-83cc-11ee-b962-0242ac120002",
		Name:             "Exercise A",
		Question:         "What is the capital of France?",
		Answers:          "[Paris, Amsterdam, Brussels, Berlin]",
		PosCorrectAnswer: 0,
		QuestionTypeID:   "multiple_choice",
		Difficulty:       2.5,
		CreatedAt:        &timestamp,
		UpdatedAt:        &timestamp,
		SoftDeleted:      &softDeleted,
	})

	r.exercises = append(r.exercises, &model.Exercise{
		ID:               "94ed3b8c-83ce-11ee-b962-0242ac120002",
		ClassID:          "7c1ee614-83ce-11ee-b962-0242ac120002",
		Name:             "Exercise A",
		Question:         "What is the capital of France?",
		Answers:          "[Paris, Amsterdam, Brussels, Berlin]",
		PosCorrectAnswer: 0,
		QuestionTypeID:   "multiple_choice",
		Difficulty:       2.5,
		CreatedAt:        &timestamp,
		UpdatedAt:        &timestamp,
		SoftDeleted:      &softDeleted,
	})

	/// RESULTS ///
	r.results = append(r.results, &model.Result{
		ID:          "fe1951f6-83cc-11ee-b962-0242ac120002",
		ExerciseID:  "af2cef6a-83ce-11ee-b962-0242ac120002",
		UserID:      "3a6085b2-83cd-11ee-b962-0242ac120002",
		ClassID:     "f2f371e4-83cc-11ee-b962-0242ac120002",
		ModuleID:    "e41189cc-83cc-11ee-b962-0242ac120002",
		Input:       "Example input",
		Result:      "Example result",
		CreatedAt:   &timestamp,
		UpdatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	})

	/// LEADERBOARDS ///
	r.leaderboard = append(r.leaderboard, &model.LeaderboardRow{
		ID:       "026b96b0-83cd-11ee-b962-0242ac120002",
		Name:     "Bram Terlouw",
		Rating:   1200,
		Position: 1,
	})
	r.leaderboard = append(r.leaderboard, &model.LeaderboardRow{
		ID:       "104acb2a-83cd-11ee-b962-0242ac120002",
		Name:     "Bastiaan van der Bijl",
		Rating:   1100,
		Position: 1,
	})
	r.leaderboard = append(r.leaderboard, &model.LeaderboardRow{
		ID:       "13fa03b2-83cd-11ee-b962-0242ac120002",
		Name:     "Merlijn Busch",
		Rating:   1000,
		Position: 1,
	})
}
