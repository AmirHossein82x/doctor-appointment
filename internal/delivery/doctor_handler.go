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

func SetUpDoctorRoutes(router *gin.RouterGroup) {
	log := logger.SetUpLogger()
	log.Info("Setting up auth routes")
	doctorRepo := repository.NewDoctorRepository()
	SmsService := infrastructure.NewKavenegarSmsService()
	var doctorHandler ports.DoctorService = services.NewDoctorService(doctorRepo, log, SmsService)

	doctorRoute := router.Group("/doctor", middleware.AuthMiddleware(constants.DoctorRole))
	doctorRoute.POST("/create-appointment", doctorHandler.CreateAppointment)
	doctorRoute.GET("/available-appointments", doctorHandler.GetAvailableAppointments)
}
