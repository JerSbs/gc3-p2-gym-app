package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetBMIFromAPI(weight, height float64) (string, error) {
	url := fmt.Sprintf("https://%s/metric?weight=%.2f&height=%.2f",
		os.Getenv("BMI_API_HOST"), weight, height)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("x-rapidapi-key", os.Getenv("BMI_API_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("BMI_API_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
