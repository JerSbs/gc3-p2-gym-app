package repository

import (
	"gc3-p2-gym-app-JerSbs/models"

	"gorm.io/gorm"
)

// Interface
type UserLoginRepository interface {
	FindByEmail(email string) (*models.User, error)
}

// Struct
type userLoginRepository struct {
	db *gorm.DB
}

// Constructor
func NewUserLoginRepository(db *gorm.DB) UserLoginRepository {
	return &userLoginRepository{db}
}

// Implementation
func (r *userLoginRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
