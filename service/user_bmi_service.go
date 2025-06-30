package service

import (
	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/models"
	"gc3-p2-gym-app-JerSbs/utils"
)

func GetUserBMIService(userID uint) (*dto.BMIResponse, error) {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return nil, ErrNotFound
	}

	// Call external API from utils
	bmiResp, err := utils.CalculateBMI(float64(user.Weight), float64(user.Height)/100) // convert cm to meter
	if err != nil {
		return nil, ErrInternal
	}

	return bmiResp, nil
}
