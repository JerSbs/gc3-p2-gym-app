package service

import (
	"errors"

	"gorm.io/gorm"

	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/models"
	"gc3-p2-gym-app-JerSbs/repository"
)

func DeleteWorkoutService(workoutID uint, userID uint) (*dto.WorkoutResponse, error) {
	db := config.GetDB()
	repo := repository.NewWorkoutDeleteRepository(db)

	workout, err := repo.GetWorkoutWithExercisesByID(workoutID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrInternal
	}

	if workout.UserID != userID {
		return nil, ErrForbidden
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		// Delete exercises first
		if len(workout.Exercises) > 0 {
			if err := tx.Where("workout_id = ?", workout.ID).Delete(&models.Exercise{}).Error; err != nil {
				return err
			}
		}

		// Then delete workout
		if err := tx.Delete(&workout).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, ErrInternal
	}

	return &dto.WorkoutResponse{
		ID:          workout.ID,
		Name:        workout.Name,
		Description: workout.Description,
		UserID:      workout.UserID,
	}, nil
}
