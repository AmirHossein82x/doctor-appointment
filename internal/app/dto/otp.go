package dto

type OTPRequest struct {
	PhoneNumber string `json:"phone_number"`
}


type VerifyOTPRequest struct {
	PhoneNumber string `json:"phone_number"`
	OTP string `json:"otp_code"`
}