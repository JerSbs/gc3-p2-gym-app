package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"p2-graded-challenge-3-JerSbs/dto"
)

func CalculateBMI(weight, height float64) (*dto.BMIResponse, error) {
	url := fmt.Sprintf("https://body-mass-index-bmi-calculator.p.rapidapi.com/metric?weight=%.2f&height=%.2f", weight, height)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-rapidapi-key", os.Getenv("RAPID_API_KEY"))
	req.Header.Add("x-rapidapi-host", "body-mass-index-bmi-calculator.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var bmiResp dto.BMIResponse
	err = json.Unmarshal(body, &bmiResp)
	if err != nil {
		return nil, err
	}

	return &bmiResp, nil
}
