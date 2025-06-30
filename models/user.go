package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"not null;unique" json:"email"`
	FullName  string    `gorm:"not null;unique" json:"full_name"`
	Password  string    `gorm:"not null" json:"-"`
	Weight    int       `gorm:"not null" json:"weight"`
	Height    int       `gorm:"not null" json:"height"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Name The Table
func (User) TableName() string {
	return "users"
}
