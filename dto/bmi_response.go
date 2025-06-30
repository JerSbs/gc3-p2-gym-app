package dto

type UserBMIResponse struct {
	ID      uint    `json:"id" example:"1"`
	Name    string  `json:"name" example:"Test"`
	Email   string  `json:"email" example:"example@example.com"`
	Weight  int     `json:"weight" example:"60"`
	Height  int     `json:"height" example:"165"`
	BMIData BMIData `json:"bmi_data"`
}

type BMIData struct {
	BMI            float64 `json:"bmi" example:"22.04"`
	Weight         string  `json:"weight" example:"60.00"`
	Height         string  `json:"height" example:"165.00"`
	WeightCategory string  `json:"weightCategory" example:"Normal Weight"`
}
