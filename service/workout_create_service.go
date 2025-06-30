package service

import (
	"p2-graded-challenge-3-JerSbs/config"
	"p2-graded-challenge-3-JerSbs/dto"
	"p2-graded-challenge-3-JerSbs/models"
	"p2-graded-challenge-3-JerSbs/repository"
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
