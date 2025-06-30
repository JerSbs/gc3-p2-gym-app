package repository

import (
	"gorm.io/gorm"
	"p2-graded-challenge-3-JerSbs/models"
)

type WorkoutDeleteRepository interface {
	GetWorkoutWithExercisesByID(id uint) (*models.Workout, error)
	DeleteWorkout(workout *models.Workout) error
}

type workoutDeleteRepository struct {
	db *gorm.DB
}

func NewWorkoutDeleteRepository(db *gorm.DB) WorkoutDeleteRepository {
	return &workoutDeleteRepository{db}
}

func (r *workoutDeleteRepository) GetWorkoutWithExercisesByID(id uint) (*models.Workout, error) {
	var workout models.Workout
	err := r.db.Preload("Exercises").First(&workout, id).Error
	return &workout, err
}

func (r *workoutDeleteRepository) DeleteWorkout(workout *models.Workout) error {
	return r.db.Delete(workout).Error
}
