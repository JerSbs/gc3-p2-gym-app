package repository

import (
	"gc3-p2-gym-app-JerSbs/models"

	"gorm.io/gorm"
)

type LogGetRepository interface {
	GetLogsByUserID(userID uint) ([]models.Log, error)
}

type logGetRepository struct {
	db *gorm.DB
}

func NewLogGetRepository(db *gorm.DB) LogGetRepository {
	return &logGetRepository{db}
}

func (r *logGetRepository) GetLogsByUserID(userID uint) ([]models.Log, error) {
	var logs []models.Log
	err := r.db.Preload("Exercise").Where("user_id = ?", userID).Find(&logs).Error
	return logs, err
}
