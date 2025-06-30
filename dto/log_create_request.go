package dto

type LogCreateRequest struct {
	ExerciseID uint   `json:"exercise_id" validate:"required"`
	Weight     int    `json:"weight" validate:"required"`
	RepCount   int    `json:"repition_count" validate:"required"`
	SetCount   int    `json:"set_count" validate:"required"`
	CreatedAt  string `json:"created_at"` // optional or for logging timestamp
}

type LogResponse struct {
	ID         uint `json:"id"`
	ExerciseID uint `json:"exercise_id"`
	UserID     uint `json:"user_id"`
	Weight     int  `json:"weight"`
	RepCount   int  `json:"repition_count"`
	SetCount   int  `json:"set_count"`
}
