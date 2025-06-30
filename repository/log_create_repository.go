package repository

import (
	"p2-graded-challenge-3-JerSbs/models"

	"gorm.io/gorm"
)

type LogCreateRepository interface {
	GetExerciseWithWorkout(id uint) (*models.Exercise, error)
	CreateLog(log *models.Log) error
}

type logCreateRepository struct {
	db *gorm.DB
}

func NewLogCreateRepository(db *gorm.DB) LogCreateRepository {
	return &logCreateRepository{db}
}

// Preload Workout → used to verify ownership by workout.user_id
func (r *logCreateRepository) GetExerciseWithWorkout(id uint) (*models.Exercise, error) {
	var exercise models.Exercise
	err := r.db.Preload("Workout").First(&exercise, id).Error
	return &exercise, err
}

func (r *logCreateRepository) CreateLog(log *models.Log) error {
	return r.db.Create(log).Error
}
