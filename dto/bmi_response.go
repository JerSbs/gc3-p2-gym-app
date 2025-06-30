package dto

type BMIResponse struct {
	BMI            float64 `json:"bmi"`
	Weight         string  `json:"weight"`
	Height         string  `json:"height"`
	WeightCategory string  `json:"weight_category"`
}
