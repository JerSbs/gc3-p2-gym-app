package service

import (
	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/models"
	"gc3-p2-gym-app-JerSbs/utils"
)

type UserBMIResult struct {
	ID      uint         `json:"id"`
	Name    string       `json:"name"`
	Email   string       `json:"email"`
	Weight  int          `json:"weight"`
	Height  int          `json:"height"`
	BMIData *dto.BMIData `json:"bmi_data"`
}

func GetUserBMIService(userID uint) (*UserBMIResult, error) {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return nil, ErrNotFound
	}

	bmiData, err := utils.GetBMIFromAPI(float64(user.Weight), float64(user.Height))
	if err != nil {
		return nil, ErrExternalAPI
	}

	result := &UserBMIResult{
		ID:      user.ID,
		Name:    user.FullName,
		Email:   user.Email,
		Weight:  user.Weight,
		Height:  user.Height,
		BMIData: bmiData,
	}
	return result, nil
}
