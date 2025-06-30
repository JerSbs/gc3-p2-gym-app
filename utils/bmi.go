package utils

import "math"

func CalculateBMI(weight, height int) float64 {
	if height == 0 {
		return 0
	}
	heightInMeters := float64(height) / 100
	bmi := float64(weight) / (heightInMeters * heightInMeters)

	// Round to 2 decimal places
	return math.Round(bmi*100) / 100
}

func GetWeightCategory(bmi float64) string {
	switch {
	case bmi < 18.5:
		return "Underweight"
	case bmi >= 18.5 && bmi < 24.9:
		return "Normal Weight"
	case bmi >= 25 && bmi < 29.9:
		return "Overweight"
	default:
		return "Obese"
	}
}
