package models

type UpdateProfileRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
