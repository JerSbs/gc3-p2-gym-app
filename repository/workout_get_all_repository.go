package repository

import (
	"gc3-p2-gym-app-JerSbs/models"

	"gorm.io/gorm"
)

type WorkoutRepository interface {
	GetWorkoutsByUserID(userID uint) ([]models.Workout, error)
	CreateWorkout(workout *models.Workout) error
}

type workoutRepository struct {
	db *gorm.DB
}

func NewWorkoutRepository(db *gorm.DB) WorkoutRepository {
	return &workoutRepository{db}
}

func (r *workoutRepository) GetWorkoutsByUserID(userID uint) ([]models.Workout, error) {
	var workouts []models.Workout
	err := r.db.Where("user_id = ?", userID).Find(&workouts).Error
	return workouts, err
}

func (r *workoutRepository) CreateWorkout(workout *models.Workout) error {
	return r.db.Create(workout).Error
}
