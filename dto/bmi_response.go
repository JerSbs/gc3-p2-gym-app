package dto

type BMIData struct {
	BMI            float64 `json:"bmi" example:"22.04"`
	WeightCategory string  `json:"weightCategory" example:"Normal Weight"`
}

type UserBMIResponse struct {
	ID      uint    `json:"id" example:"2"`
	Name    string  `json:"name" example:"Jane Test"`
	Email   string  `json:"email" example:"jane@example.com"`
	Weight  int     `json:"weight" example:"60"`
	Height  int     `json:"height" example:"165"`
	BMIData BMIData `json:"bmi_data"`
}
