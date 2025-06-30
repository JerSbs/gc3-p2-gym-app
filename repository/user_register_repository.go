package repository

import (
	"gorm.io/gorm"
	"p2-graded-challenge-3-JerSbs/models"
)

type UserRegisterRepository interface {
	Create(user *models.User) error
	FindByEmailOrFullName(email, fullName string) (*models.User, error)
}

type userRegisterRepository struct {
	db *gorm.DB
}

func NewUserRegisterRepository(db *gorm.DB) UserRegisterRepository {
	return &userRegisterRepository{db}
}

func (r *userRegisterRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRegisterRepository) FindByEmailOrFullName(email, fullName string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ? OR full_name = ?", email, fullName).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
