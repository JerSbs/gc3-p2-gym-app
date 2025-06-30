package service

import (
	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/repository"
	"gc3-p2-gym-app-JerSbs/utils"
)

func LoginUser(input dto.LoginRequest) (string, error) {
	if input.Email == "" || input.Password == "" {
		return "", ErrInvalidInput
	}

	db := config.GetDB()
	userRepo := repository.NewUserLoginRepository(db)

	user, err := userRepo.FindByEmail(input.Email)
	if err != nil {
		return "", ErrUserNotFound
	}

	if !utils.CheckPasswordHash(user.Password, input.Password) {
		return "", ErrWrongPassword
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", ErrInternal
	}

	return token, nil
}
