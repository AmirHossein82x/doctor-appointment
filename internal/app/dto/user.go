package dto

import "github.com/google/uuid"

type UserRegisterRequest struct {
	Name          string `json:"name"`
	Password      string `json:"password"`
	VerifiedToken string `json:"verified_token"`
}

type UserRegisterResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	CreatedAt   string    `json:"created_at"`
}
