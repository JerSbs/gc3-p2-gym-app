package service

import (
	"errors"

	"gorm.io/gorm"
	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/repository"
)

func DeleteExerciseService(userID uint, exerciseID uint) (*dto.ExerciseResponse, error) {
	db := config.GetDB()
	repo := repository.NewExerciseDeleteRepository(db)

	exercise, err := repo.GetExerciseWithWorkout(exerciseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrInternal
	}

	if exercise.Workout.UserID != userID {
		return nil, ErrForbidden
	}

	if err := repo.DeleteExerciseAndLogs(exerciseID); err != nil {
		return nil, ErrInternal
	}

	return &dto.ExerciseResponse{
		ID:          exercise.ID,
		WorkoutID:   exercise.WorkoutID,
		Name:        exercise.Name,
		Description: exercise.Description,
	}, nil
}
