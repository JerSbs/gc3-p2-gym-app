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
	// ✅ Step 1: Call /metric
	metricURL := fmt.Sprintf("https://%s/metric?weight=%.2f&height=%.2f",
		os.Getenv("BMI_API_HOST"), weight, height)

	req, err := http.NewRequest("GET", metricURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-rapidapi-key", os.Getenv("BMI_API_KEY"))
	req.Header.Set("x-rapidapi-host", os.Getenv("BMI_API_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, _ := io.ReadAll(res.Body)
	fmt.Println("📦 /metric raw response:", string(bodyBytes))

	var raw struct {
		BMI    string `json:"bmi"`
		Weight string `json:"weight"`
		Height string `json:"height"`
	}
	if err := json.Unmarshal(bodyBytes, &raw); err != nil {
		return nil, err
	}

	bmiFloat, err := strconv.ParseFloat(raw.BMI, 64)
	if err != nil {
		return nil, err
	}

	// ✅ Step 2: Call /weight-category?bmi=<bmi_string>
	categoryURL := fmt.Sprintf("https://%s/weight-category?bmi=%s",
		os.Getenv("BMI_API_HOST"), raw.BMI)

	req2, err := http.NewRequest("GET", categoryURL, nil)
	if err != nil {
		return nil, err
	}
	req2.Header.Set("x-rapidapi-key", os.Getenv("BMI_API_KEY"))
	req2.Header.Set("x-rapidapi-host", os.Getenv("BMI_API_HOST"))

	res2, err := http.DefaultClient.Do(req2)
	if err != nil {
		return nil, err
	}
	defer res2.Body.Close()

	bodyBytes2, _ := io.ReadAll(res2.Body)
	fmt.Println("📦 /weight-category raw response:", string(bodyBytes2))

	var categoryResp struct {
		WeightCategory string `json:"weightCategory"`
	}
	if err := json.Unmarshal(bodyBytes2, &categoryResp); err != nil {
		return nil, err
	}

	return &dto.BMIData{
		BMI:            bmiFloat,
		Weight:         raw.Weight,
		Height:         raw.Height,
		WeightCategory: categoryResp.WeightCategory,
	}, nil
}
