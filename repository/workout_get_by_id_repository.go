package repository

import (
	"p2-graded-challenge-3-JerSbs/models"

	"gorm.io/gorm"
)

// Define interface
type WorkoutGetByIDRepository interface {
	GetWorkoutByID(workoutID uint) (*models.Workout, error)
}

// Struct to hold DB instance
type workoutGetByIDRepository struct {
	db *gorm.DB
}

// Constructor
func NewWorkoutGetByIDRepository(db *gorm.DB) WorkoutGetByIDRepository {
	return &workoutGetByIDRepository{db}
}

// Method to fetch workout by ID, with Exercises
func (r *workoutGetByIDRepository) GetWorkoutByID(workoutID uint) (*models.Workout, error) {
	var workout models.Workout
	err := r.db.Preload("Exercises").First(&workout, workoutID).Error
	if err != nil {
		return nil, err
	}
	return &workout, nil
}
