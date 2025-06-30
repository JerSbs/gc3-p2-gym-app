package service

import (
	"errors"

	"gorm.io/gorm"

	"p2-graded-challenge-3-JerSbs/config"
	"p2-graded-challenge-3-JerSbs/dto"
	"p2-graded-challenge-3-JerSbs/models"
	"p2-graded-challenge-3-JerSbs/repository"
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
