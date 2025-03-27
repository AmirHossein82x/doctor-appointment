package dto

import "github.com/google/uuid"

type UserRegisterRequest struct {
	Name          string `json:"name"`
	Password      string `json:"password"`
	VerifiedToken string `json:"verified_token"`
}

type UserRegisterResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	PhoneNumber  string    `json:"phone_number"`
	Role         string    `json:"role"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    string    `json:"created_at"`
}

type UserLoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type ForgetPasswordRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type PasswordResetRequest struct {
	Password       string `json:"password"`
	PasswordRetype string `json:"password_retype"`
}
