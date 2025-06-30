package service

import (
	"p2-graded-challenge-3-JerSbs/config"
	"p2-graded-challenge-3-JerSbs/dto"
	"p2-graded-challenge-3-JerSbs/models"
	"p2-graded-challenge-3-JerSbs/repository"
	"p2-graded-challenge-3-JerSbs/utils"
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
