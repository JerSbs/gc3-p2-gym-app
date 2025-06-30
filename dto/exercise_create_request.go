package dto

type ExerciseCreateRequest struct {
	WorkoutID   uint   `json:"workout_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type ExerciseResponse struct {
	ID          uint   `json:"id"`
	WorkoutID   uint   `json:"workout_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
