package delivery

import (
	"github.com/AmirHossein82x/doctor-appointment/internal/app/services"
	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
	"github.com/AmirHossein82x/doctor-appointment/internal/logger"
	"github.com/AmirHossein82x/doctor-appointment/internal/repository"
	"github.com/gin-gonic/gin"
)


func SetUpUserRoutes(router *gin.RouterGroup) {
	log := logger.SetUpLogger()
	log.Info("Setting up User routes")
	userRepo := repository.NewUserRepository(infrastructure.GetRedisClient())
	SmsService := infrastructure.NewKavenegarSmsService()
	UserHandler := services.NewUserService(userRepo, log, SmsService)

	userRoute := router.Group("/users")
	userRoute.POST("/register", UserHandler.Register)
	userRoute.POST("/login", UserHandler.Login)
	userRoute.POST("/verify-access-token", UserHandler.VerifyAccessToken)
	userRoute.POST("/get-access-token-by-refresh-token", UserHandler.GetAccessTokenByRefreshToken)
	userRoute.POST("/forget-password", UserHandler.ForgetPassword)
	userRoute.POST("/reset-password", UserHandler.ResetPassword)
}