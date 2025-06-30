package dto

type UserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Weight   int    `json:"weight"`
	Height   int    `json:"height"`
}
