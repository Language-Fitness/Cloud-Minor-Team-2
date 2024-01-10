package responses

var CreateExerciseResponse struct {
	CreateExercise struct {
		ID       string
		ClassID  string `json:"class_Id"`
		ModuleID string `json:"module_id"`
		Name     string
		Question string
		Answers  []struct {
			Value   string
			Correct bool
		}
		Difficulty string
		MadeBy     string `json:"made_by"`
	}
}
