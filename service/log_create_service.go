package service

import (
	"errors"

	"gorm.io/gorm"

	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/models"
	"gc3-p2-gym-app-JerSbs/repository"
)

func CreateLogService(userID uint, input dto.LogCreateRequest) (*dto.LogResponse, error) {
	db := config.GetDB()
	repo := repository.NewLogCreateRepository(db)

	exercise, err := repo.GetExerciseWithWorkout(input.ExerciseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrInternal
	}

	if exercise.Workout.UserID != userID {
		return nil, ErrForbidden
	}

	log := models.Log{
		ExerciseID: input.ExerciseID,
		UserID:     userID,
		Weight:     input.Weight,
		RepCount:   input.RepCount,
		SetCount:   input.SetCount,
	}

	if err := repo.CreateLog(&log); err != nil {
		return nil, ErrInternal
	}

	return &dto.LogResponse{
		ID:         log.ID,
		ExerciseID: log.ExerciseID,
		UserID:     log.UserID,
		Weight:     log.Weight,
		RepCount:   log.RepCount,
		SetCount:   log.SetCount,
	}, nil
}
