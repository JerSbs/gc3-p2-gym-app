package service

import (
	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/models"
	"gc3-p2-gym-app-JerSbs/repository"
)

func CreateWorkoutService(userID uint, input dto.WorkoutCreateRequest) (*models.Workout, error) {
	if input.Name == "" || input.Description == "" {
		return nil, ErrInvalidInput
	}

	db := config.GetDB()
	repo := repository.NewWorkoutCreateRepository(db)

	workout := &models.Workout{
		Name:        input.Name,
		Description: input.Description,
		UserID:      userID,
	}

	if err := repo.CreateWorkout(workout); err != nil {
		return nil, ErrInternal
	}

	return workout, nil
}
