package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gc3-p2-gym-app-JerSbs/dto"
)

func GetBMIFromAPI(weight, height float64) (*dto.BMIData, error) {
	url := fmt.Sprintf("https://%s/metric?weight=%.2f&height=%.2f",
		os.Getenv("BMI_API_HOST"), weight, height)

	req, err := http.NewRequest("GET", url, nil)
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

	var bmi dto.BMIData
	err = json.NewDecoder(res.Body).Decode(&bmi)
	if err != nil {
		return nil, err
	}

	return &bmi, nil
}
