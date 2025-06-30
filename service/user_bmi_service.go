package service

import (
	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/models"
	"gc3-p2-gym-app-JerSbs/utils"
)

func GetUserBMIService(userID uint) (*dto.UserBMIResponse, error) {
	db := config.GetDB()
	var user models.User

	if err := db.First(&user, userID).Error; err != nil {
		return nil, ErrNotFound
	}

	bmiData, err := utils.GetBMIFromAPI(float64(user.Weight), float64(user.Height))
	if err != nil {
		return nil, err
	}

	return &dto.UserBMIResponse{
		ID:      user.ID,
		Name:    user.FullName,
		Email:   user.Email,
		Weight:  user.Weight,
		Height:  user.Height,
		BMIData: *bmiData,
	}, nil
}
