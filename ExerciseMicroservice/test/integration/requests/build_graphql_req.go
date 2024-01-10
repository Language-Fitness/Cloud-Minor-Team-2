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
