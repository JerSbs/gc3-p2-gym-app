package dto

type WorkoutDetailResponse struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Exercises   []ExerciseItem `json:"exercises"`
}
