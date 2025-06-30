package models

import "gorm.io/gorm"

type Workout struct {
	gorm.Model
	Name        string     `gorm:"not null" json:"name"`
	Description string     `gorm:"not null" json:"description"`
	UserID      uint       `gorm:"not null" json:"user_id"`
	Exercises   []Exercise `gorm:"foreignKey:WorkoutID" json:"exercises"`
}
