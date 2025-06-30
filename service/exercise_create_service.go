package service

import (
	"errors"

	"gorm.io/gorm"

	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/models"
	"gc3-p2-gym-app-JerSbs/repository"
)

func CreateExerciseService(userID uint, input dto.ExerciseCreateRequest) (*dto.ExerciseResponse, error) {
	db := config.GetDB()
	repo := repository.NewExerciseCreateRepository(db)

	workout, err := repo.GetWorkoutByID(input.WorkoutID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrInternal
	}

	if workout.UserID != userID {
		return nil, ErrForbidden
	}

	exercise := models.Exercise{
		WorkoutID:   input.WorkoutID,
		Name:        input.Name,
		Description: input.Description,
	}

	if err := repo.CreateExercise(&exercise); err != nil {
		return nil, ErrInternal
	}

	return &dto.ExerciseResponse{
		ID:          exercise.ID,
		WorkoutID:   exercise.WorkoutID,
		Name:        exercise.Name,
		Description: exercise.Description,
	}, nil
}
