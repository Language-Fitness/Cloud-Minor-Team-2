package requests

var (
	ClassID    = "8e45c3b9-1cbb-4024-8ce7-ddee19940188"
	ModuleID   = "8e45c3b9-1cbb-4024-8ce7-ddee19940188"
	Name       = "Grammar, module 1, exercise 1"
	Question   = "What is the answer?"
	Difficulty = "A1"
	Answers    = []map[string]interface{}{
		{"value": "Option A", "correct": true},
		{"value": "Option B", "correct": false},
	}
)

var CreateExerciseMutation = `
    mutation($exerciseInput: ExerciseInput!) {
        CreateExercise(exercise: $exerciseInput) {
            id
            class_Id
            module_id
            name
            question
            answers {
                value
                correct
            }
            difficulty
            made_by
        }
    }`

func GenerateExerciseInput() map[string]interface{} {
	return map[string]interface{}{
		"class_Id":   ClassID,
		"module_id":  ModuleID,
		"name":       Name,
		"question":   Question,
		"difficulty": Difficulty,
		"answers":    Answers,
	}
}

func GenerateExerciseInputInvalidClassId() map[string]interface{} {
	// Set an invalid UUID for ClassID
	return map[string]interface{}{
		"class_Id":   "56565656b", // This is an invalid UUID
		"module_id":  ModuleID,
		"name":       Name,
		"question":   Question,
		"difficulty": Difficulty,
		"answers":    Answers,
	}
}

func GenerateExerciseInputInvalidModuleId() map[string]interface{} {
	// Set an invalid UUID for ModuleID
	return map[string]interface{}{
		"class_Id":   ClassID,
		"module_id":  "56565656b", // This is an invalid UUID
		"name":       Name,
		"question":   Question,
		"difficulty": Difficulty,
		"answers":    Answers,
	}
}

func GenerateExerciseInputInvalidName() map[string]interface{} {
	// Set an invalid name
	return map[string]interface{}{
		"class_Id":   ClassID,
		"module_id":  ModuleID,
		"name":       "New Exercise2222222222222222222222222222222222222222222222222222222222222222222222222222",
		"question":   Question,
		"difficulty": Difficulty,
		"answers":    Answers,
	}
}

func GenerateExerciseInputInvalidQuestion() map[string]interface{} {
	// Set an invalid question
	return map[string]interface{}{
		"class_Id":   ClassID,
		"module_id":  ModuleID,
		"name":       Name,
		"question":   "What is the question?0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		"difficulty": Difficulty,
		"answers":    Answers,
	}
}

func GenerateExerciseInputNoCorrectAnswers() map[string]interface{} {
	// Set no correct answers
	return map[string]interface{}{
		"class_Id":   ClassID,
		"module_id":  ModuleID,
		"name":       Name,
		"question":   Question,
		"difficulty": Difficulty,
		"answers": []map[string]interface{}{
			{"value": "Option A", "correct": false},
			{"value": "Option B", "correct": false},
		},
	}
}

func GenerateExerciseInputNoAnswers() map[string]interface{} {
	// Set no answers
	return map[string]interface{}{
		"class_Id":   ClassID,
		"module_id":  ModuleID,
		"name":       Name,
		"question":   Question,
		"difficulty": Difficulty,
		"answers":    []map[string]interface{}{},
	}
}

func GenerateExerciseInputNoIncorrectAnswers() map[string]interface{} {
	// Set no incorrect answers
	return map[string]interface{}{
		"class_Id":   ClassID,
		"module_id":  ModuleID,
		"name":       Name,
		"question":   Question,
		"difficulty": Difficulty,
		"answers": []map[string]interface{}{
			{"value": "Option A", "correct": true},
			{"value": "Option B", "correct": true},
		},
	}
}
