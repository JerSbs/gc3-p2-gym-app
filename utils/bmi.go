package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"gc3-p2-gym-app-JerSbs/dto"
)

func GetBMIFromAPI(weight, height float64) (*dto.BMIData, error) {
	// Step 1: Call /metric?weight=...&height=...
	metricURL := fmt.Sprintf("https://%s/metric?weight=%.2f&height=%.2f",
		os.Getenv("BMI_API_HOST"), weight, height)

	req, err := http.NewRequest("GET", metricURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-key", os.Getenv("BMI_API_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("BMI_API_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// decode /metric response
	var raw struct {
		BMI    string `json:"bmi"`
		Weight string `json:"weight"`
		Height string `json:"height"`
	}
	if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
		return nil, err
	}

	// convert BMI string to float64
	bmiValue, err := strconv.ParseFloat(raw.BMI, 64)
	if err != nil {
		return nil, err
	}

	// Step 2: Call /weight-category?bmi=<bmi string>
	categoryURL := fmt.Sprintf("https://%s/weight-category?bmi=%s",
		os.Getenv("BMI_API_HOST"), raw.BMI)

	req2, err := http.NewRequest("GET", categoryURL, nil)
	if err != nil {
		return nil, err
	}
	req2.Header.Add("x-rapidapi-key", os.Getenv("BMI_API_KEY"))
	req2.Header.Add("x-rapidapi-host", os.Getenv("BMI_API_HOST"))

	res2, err := http.DefaultClient.Do(req2)
	if err != nil {
		return nil, err
	}
	defer res2.Body.Close()

	// decode /weight-category response
	var categoryResp struct {
		WeightCategory string `json:"weightCategory"`
	}
	if err := json.NewDecoder(res2.Body).Decode(&categoryResp); err != nil {
		return nil, err
	}

	// Combine all into final result
	result := &dto.BMIData{
		BMI:            bmiValue,
		Weight:         raw.Weight,
		Height:         raw.Height,
		WeightCategory: categoryResp.WeightCategory,
	}

	return result, nil
}
