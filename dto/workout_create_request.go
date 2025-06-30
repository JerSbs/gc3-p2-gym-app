package dto

type WorkoutCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
