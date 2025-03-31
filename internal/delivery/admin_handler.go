package delivery

import (
	"github.com/AmirHossein82x/doctor-appointment/internal/app/constants"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/ports"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/services"
	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
	"github.com/AmirHossein82x/doctor-appointment/internal/logger"
	"github.com/AmirHossein82x/doctor-appointment/internal/middleware"
	"github.com/AmirHossein82x/doctor-appointment/internal/repository"
	"github.com/gin-gonic/gin"
)

func SetUpAdminRoutes(router *gin.RouterGroup) {
	log := logger.SetUpLogger()
	log.Info("Setting up Admin routes")
	AdminRepository := repository.NewAdminRepository()
	SmsService := infrastructure.NewKavenegarSmsService()
	var AdminHandler ports.AdminService = services.NewAdminService(AdminRepository, log, SmsService)

	AdminRoute := router.Group("/admin", middleware.AuthMiddleware(constants.AdminRole))

	AdminRoute.GET("/get-all-users", AdminHandler.GetAllUsers)
}
