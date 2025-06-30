package dto

type BMIData struct {
	BMI string `json:"bmi"`
}

type UserBMIResponse struct {
	ID      uint    `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Weight  int     `json:"weight"`
	Height  int     `json:"height"`
	BMIData BMIData `json:"bmi_data"`
}
