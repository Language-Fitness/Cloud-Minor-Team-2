package entity

type ExerciseType struct {
	ID          int              `json:"id"`
	Settings    ExerciseSettings `json:"exercise_settings"`
	CreatedAt   string           `json:"created_at"`
	UpdatedAt   string           `json:"updated_at"`
	SoftDeleted string           `json:"soft_deleted"`
}
