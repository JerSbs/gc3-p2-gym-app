package service

import (
	"encoding/json"
	"fmt"
	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/models"
	"io"
	"net/http"
	"os"
)

func GetUserBMIService(userID uint) (*dto.UserBMIResponse, error) {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	url := fmt.Sprintf("https://body-mass-index-bmi-calculator.p.rapidapi.com/metric?weight=%d&height=%d", user.Weight, user.Height)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-rapidapi-key", os.Getenv("BMI_API_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("BMI_API_HOST"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("📦 /metric raw response: %s\n", string(body))

	var result struct {
		BMI float64 `json:"bmi"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	response := &dto.UserBMIResponse{
		ID:      user.ID,
		Name:    user.FullName,
		Email:   user.Email,
		Weight:  user.Weight,
		Height:  user.Height,
		BMIData: dto.BMIData{BMI: result.BMI},
	}

	return response, nil
}
