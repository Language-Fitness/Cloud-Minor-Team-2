package entity

type ExerciseResult struct {
	ID          int    `json:"id"`
	ExerciseID  int    `json:"exercise_id"`
	UserID      int    `json:"user_id"`
	Input       string `json:"input"`
	Result      bool   `json:"result"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	SoftDeleted string `json:"soft_deleted"`
}
