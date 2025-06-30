package repository

import (
	"gorm.io/gorm"
	"gc3-p2-gym-app-JerSbs/models"
)

type WorkoutCreateRepository interface {
	CreateWorkout(workout *models.Workout) error
}

type workoutCreateRepository struct {
	db *gorm.DB
}

func NewWorkoutCreateRepository(db *gorm.DB) WorkoutCreateRepository {
	return &workoutCreateRepository{db}
}

func (r *workoutCreateRepository) CreateWorkout(workout *models.Workout) error {
	return r.db.Create(workout).Error
}
