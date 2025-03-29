package ports

import "github.com/gin-gonic/gin"

type OtpService interface {
	GenerateOTP(*gin.Context)
	VerifyOTP(*gin.Context)
}

type OtpRepository interface {
	GenerateOTP(string) (int, error)
	VerifyOTP(string, string) (bool, error)
	GenerateVerifyOtpToken(string) (string, error)
}
