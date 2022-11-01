package dto

type UserResponse struct {
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
