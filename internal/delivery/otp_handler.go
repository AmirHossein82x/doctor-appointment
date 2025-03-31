package delivery

import (
	"github.com/AmirHossein82x/doctor-appointment/internal/app/ports"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/services"
	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
	"github.com/AmirHossein82x/doctor-appointment/internal/logger"
	"github.com/AmirHossein82x/doctor-appointment/internal/repository"
	"github.com/gin-gonic/gin"
)


func SetUpOtpRoutes(router *gin.RouterGroup) {
	log := logger.SetUpLogger()
	log.Info("Setting up OTP routes")
	otpRep := repository.NewOtpRepo(infrastructure.GetRedisClient())
	SmsService := infrastructure.NewKavenegarSmsService()
	var OtpHandler ports.OtpService = services.NewOTPService(otpRep, log, SmsService)

	otpRoute := router.Group("/otp")

	otpRoute.POST("/generate", OtpHandler.GenerateOTP)
	otpRoute.POST("/verify", OtpHandler.VerifyOTP)
}
