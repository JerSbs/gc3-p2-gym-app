package repository

import (
	"gorm.io/gorm"
	"p2-graded-challenge-3-JerSbs/models"
)

type WorkoutUpdateRepository interface {
	GetWorkoutByID(id uint) (*models.Workout, error)
	UpdateWorkout(workout *models.Workout) error
}

type workoutUpdateRepository struct {
	db *gorm.DB
}

func NewWorkoutUpdateRepository(db *gorm.DB) WorkoutUpdateRepository {
	return &workoutUpdateRepository{db}
}

func (r *workoutUpdateRepository) GetWorkoutByID(id uint) (*models.Workout, error) {
	var workout models.Workout
	err := r.db.First(&workout, id).Error
	return &workout, err
}

func (r *workoutUpdateRepository) UpdateWorkout(workout *models.Workout) error {
	return r.db.Save(workout).Error
}
