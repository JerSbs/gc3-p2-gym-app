package service

import (
	"p2-graded-challenge-3-JerSbs/config"
	"p2-graded-challenge-3-JerSbs/models"
	"p2-graded-challenge-3-JerSbs/repository"
)

func GetUserProfile(userID uint) (*models.User, error) {
	repo := repository.NewUserProfileRepository(config.GetDB())
	user, err := repo.FindByID(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
