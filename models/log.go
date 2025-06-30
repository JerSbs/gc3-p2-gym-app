package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	ExerciseID uint `gorm:"not null" json:"exercise_id"`
	UserID     uint `gorm:"not null" json:"user_id"`
	SetCount   int  `gorm:"not null" json:"set_count"`
	RepCount   int  `gorm:"not null" json:"repition_count"`
	Weight     int  `gorm:"not null" json:"weight"`

	Exercise Exercise `gorm:"foreignKey:ExerciseID" json:"-"`
}
