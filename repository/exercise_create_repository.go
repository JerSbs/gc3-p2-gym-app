package repository

import (
	"p2-graded-challenge-3-JerSbs/models"

	"gorm.io/gorm"
)

type ExerciseCreateRepository interface {
	GetWorkoutByID(id uint) (*models.Workout, error)
	CreateExercise(exercise *models.Exercise) error
}

type exerciseCreateRepository struct {
	db *gorm.DB
}

func NewExerciseCreateRepository(db *gorm.DB) ExerciseCreateRepository {
	return &exerciseCreateRepository{db}
}

func (r *exerciseCreateRepository) GetWorkoutByID(id uint) (*models.Workout, error) {
	var workout models.Workout
	err := r.db.First(&workout, id).Error
	return &workout, err
}

func (r *exerciseCreateRepository) CreateExercise(exercise *models.Exercise) error {
	return r.db.Create(exercise).Error
}
