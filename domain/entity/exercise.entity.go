package entity

type Exercise struct {
	ID             int      `json:"id"`
	CourseID       int      `json:"course_id"`
	Name           string   `json:"name"`
	Question       string   `json:"question"`
	Answer         string   `json:"answer"`
	QuestionTypeID int      `json:"question_type_id"`
	Tags           []string `json:"tags"`
	Difficulty     int      `json:"difficulty"`
	MadeBy         string   `json:"made_by"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
	SoftDeleted    string   `json:"soft_deleted"`
}
