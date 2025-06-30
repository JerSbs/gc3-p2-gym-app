package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"gc3-p2-gym-app-JerSbs/dto"
)

func GetBMIFromAPI(weight, height float64) (*dto.BMIData, error) {
	// Step 1: Call /metric
	metricURL := fmt.Sprintf("https://%s/metric?weight=%.2f&height=%.2f",
		os.Getenv("BMI_API_HOST"), weight, height)

	req, _ := http.NewRequest("GET", metricURL, nil)
	req.Header.Add("x-rapidapi-key", os.Getenv("BMI_API_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("BMI_API_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var raw struct {
		BMI string `json:"bmi"`
	}
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	bmiValue, _ := strconv.ParseFloat(raw.BMI, 64)

	// Step 2: Call /weight-category?bmi=<string>
	categoryURL := fmt.Sprintf("https://%s/weight-category?bmi=%s",
		os.Getenv("BMI_API_HOST"), raw.BMI)

	req2, _ := http.NewRequest("GET", categoryURL, nil)
	req2.Header.Add("x-rapidapi-key", os.Getenv("BMI_API_KEY"))
	req2.Header.Add("x-rapidapi-host", os.Getenv("BMI_API_HOST"))

	res2, err := http.DefaultClient.Do(req2)
	if err != nil {
		return nil, err
	}
	defer res2.Body.Close()

	body2, _ := io.ReadAll(res2.Body)
	var result struct {
		WeightCategory string `json:"weightCategory"`
	}
	if err := json.Unmarshal(body2, &result); err != nil {
		return nil, err
	}

	return &dto.BMIData{
		BMI:            bmiValue,
		WeightCategory: result.WeightCategory,
	}, nil
}
