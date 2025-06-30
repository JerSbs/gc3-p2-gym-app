package service

import (
	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/repository"
)

func GetAllWorkoutsService(userID uint) ([]dto.WorkoutResponse, error) {
	db := config.GetDB()
	repo := repository.NewWorkoutRepository(db)

	workouts, err := repo.GetWorkoutsByUserID(userID)
	if err != nil {
		return nil, ErrInternal
	}

	// Convert to DTO
	var response []dto.WorkoutResponse
	for _, w := range workouts {
		response = append(response, dto.WorkoutResponse{
			ID:          w.ID,
			Name:        w.Name,
			Description: w.Description,
			UserID:      w.UserID,
		})
	}

	return response, nil
}
