package entity

type ExerciseSettings struct {
	Timer         bool `json:"timer"`
	TimeInSeconds int  `json:"time_in_seconds"`
}
