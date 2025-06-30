package service

import (
	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/models"
	"gc3-p2-gym-app-JerSbs/repository"
)

func GetUserProfile(userID uint) (*models.User, error) {
	repo := repository.NewUserProfileRepository(config.GetDB())
	user, err := repo.FindByID(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
