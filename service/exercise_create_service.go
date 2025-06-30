package service

import (
	"errors"

	"gorm.io/gorm"

	"p2-graded-challenge-3-JerSbs/config"
	"p2-graded-challenge-3-JerSbs/dto"
	"p2-graded-challenge-3-JerSbs/models"
	"p2-graded-challenge-3-JerSbs/repository"
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
