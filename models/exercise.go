package models

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	WorkoutID   uint   `gorm:"not null" json:"workout_id"`

	Workout Workout `gorm:"foreignKey:WorkoutID" json:"-"`
}
