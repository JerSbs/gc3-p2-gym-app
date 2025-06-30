package service

import (
	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/models"
	"gc3-p2-gym-app-JerSbs/repository"
	"gc3-p2-gym-app-JerSbs/utils"
)

func RegisterUser(input dto.RegisterRequest) (*models.User, error) {
	if input.Email == "" || input.Password == "" || input.FullName == "" {
		return nil, ErrInvalidInput
	}

	db := config.GetDB()
	userRepo := repository.NewUserRegisterRepository(db)

	_, err := userRepo.FindByEmailOrFullName(input.Email, input.FullName)
	if err == nil {
		return nil, ErrEmailAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, ErrInternal
	}

	user := models.User{
		Email:    input.Email,
		FullName: input.FullName,
		Password: hashedPassword,
		Weight:   input.Weight,
		Height:   input.Height,
	}

	if err := userRepo.Create(&user); err != nil {
		return nil, ErrInternal
	}

	return &user, nil
}
