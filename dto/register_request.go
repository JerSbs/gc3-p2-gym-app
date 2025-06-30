package dto

type RegisterRequest struct {
	Email    string `json:"email" form:"email"`
	FullName string `json:"full_name" form:"full_name"`
	Password string `json:"password" form:"password"`
	Weight   int    `json:"weight" form:"weight"`
	Height   int    `json:"height" form:"height"`
}
