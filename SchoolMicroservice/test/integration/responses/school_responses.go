package responses

var CreateSchoolResponse struct {
	CreateSchool struct {
		ID              string
		Name            string
		Location        string
		MadeBy          string `json:"made_by"`
		HasOpenaiAccess bool   `json:"has_openai_access"`
		OpenaiKey       string `json:"openai_key"`
		JoinCode        string `json:"join_code"`
	}
}

var GetSchoolResponse struct {
	GetSchool struct {
		ID       string
		Name     string
		Location string
		MadeBy   string `json:"made_by"`
	}
}

var UpdateSchoolResponse struct {
	UpdateSchool struct {
		ID              string
		Name            string
		Location        string
		MadeBy          string `json:"made_by"`
		HasOpenaiAccess bool   `json:"has_openai_access"`
		OpenaiKey       string `json:"openai_key"`
		JoinCode        string `json:"join_code"`
	}
}
