package service

import (
	"errors"

	"gorm.io/gorm"

	"p2-graded-challenge-3-JerSbs/config"
	"p2-graded-challenge-3-JerSbs/dto"
	"p2-graded-challenge-3-JerSbs/repository"
)

func UpdateWorkoutService(workoutID uint, userID uint, payload dto.WorkoutUpdateRequest) (*dto.WorkoutResponse, error) {
	db := config.GetDB()
	repo := repository.NewWorkoutUpdateRepository(db)

	workout, err := repo.GetWorkoutByID(workoutID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrInternal
	}

	if workout.UserID != userID {
		return nil, ErrForbidden
	}

	workout.Name = payload.Name
	workout.Description = payload.Description

	if err := repo.UpdateWorkout(workout); err != nil {
		return nil, ErrInternal
	}

	return &dto.WorkoutResponse{
		ID:          workout.ID,
		Name:        workout.Name,
		Description: workout.Description,
		UserID:      workout.UserID,
	}, nil
}
