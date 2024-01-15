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

var UpdateExerciseResponse struct {
	UpdateExercise struct {
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

var GetExerciseResponse struct {
	GetExercise struct {
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

var ListExerciseResponse struct {
	ListExercise []struct {
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
