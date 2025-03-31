package delivery

import (
	"github.com/AmirHossein82x/doctor-appointment/internal/app/ports"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/services"
	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
	"github.com/AmirHossein82x/doctor-appointment/internal/logger"
	"github.com/AmirHossein82x/doctor-appointment/internal/repository"
	"github.com/gin-gonic/gin"
)

func SetUpUserRoutes(router *gin.RouterGroup) {
	log := logger.SetUpLogger()
	log.Info("Setting up auth routes")
	authRepo := repository.NewAuthRepository(infrastructure.GetRedisClient())
	SmsService := infrastructure.NewKavenegarSmsService()
	var AuthHandler ports.AuthService = services.NewAuthService(authRepo, log, SmsService)

	userRoute := router.Group("/auth")
	userRoute.POST("/register", AuthHandler.Register)
	userRoute.POST("/login", AuthHandler.Login)
	userRoute.POST("/verify-access-token", AuthHandler.VerifyAccessToken)
	userRoute.POST("/get-access-token-by-refresh-token", AuthHandler.GetAccessTokenByRefreshToken)
	userRoute.POST("/forget-password", AuthHandler.ForgetPassword)
	userRoute.POST("/reset-password", AuthHandler.ResetPassword)
}
