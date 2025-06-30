package repository

import (
	"gorm.io/gorm"
	"gc3-p2-gym-app-JerSbs/models"
)

type UserProfileRepository interface {
	FindByID(id uint) (*models.User, error)
}

type userProfileRepository struct {
	db *gorm.DB
}

func NewUserProfileRepository(db *gorm.DB) UserProfileRepository {
	return &userProfileRepository{db}
}

func (r *userProfileRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
