package dto

type LogWithExerciseResponse struct {
	ID         uint         `json:"id"`
	ExerciseID uint         `json:"exercise_id"`
	UserID     uint         `json:"user_id"`
	Weight     int          `json:"weight"`
	RepCount   int          `json:"repition_count"`
	SetCount   int          `json:"set_count"`
	Exercise   ExerciseItem `json:"exercise"`
}
