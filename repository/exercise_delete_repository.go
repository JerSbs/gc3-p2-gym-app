package repository

import (
	"gorm.io/gorm"
	"p2-graded-challenge-3-JerSbs/models"
)

type ExerciseDeleteRepository interface {
	GetExerciseWithWorkout(id uint) (*models.Exercise, error)
	DeleteExerciseAndLogs(exerciseID uint) error
}

type exerciseDeleteRepository struct {
	db *gorm.DB
}

func NewExerciseDeleteRepository(db *gorm.DB) ExerciseDeleteRepository {
	return &exerciseDeleteRepository{db}
}

func (r *exerciseDeleteRepository) GetExerciseWithWorkout(id uint) (*models.Exercise, error) {
	var exercise models.Exercise
	err := r.db.Preload("Workout").First(&exercise, id).Error
	return &exercise, err
}

func (r *exerciseDeleteRepository) DeleteExerciseAndLogs(exerciseID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Delete logs first
		if err := tx.Where("exercise_id = ?", exerciseID).Delete(&models.Log{}).Error; err != nil {
			return err
		}
		// Then delete exercise
		if err := tx.Delete(&models.Exercise{}, exerciseID).Error; err != nil {
			return err
		}
		return nil
	})
}
