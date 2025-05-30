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

func SetUpAppointmentRoutes(router *gin.RouterGroup) {
	log := logger.SetUpLogger()
	log.Info("Setting up Appointment routes")
	appointmentRepository := repository.NewAppointmentRepository()
	SmsService := infrastructure.NewKavenegarSmsService()
	var appointmentHandler ports.AppointmentService = services.NewAppointmentService(appointmentRepository, log, SmsService)
	appointmentRoute := router.Group("/appointment")
	appointmentRoute.GET("/get-doctor-profiles", appointmentHandler.GetDoctorProfiles)
	appointmentRoute.GET("/get-specialities", appointmentHandler.RetrieveSpeciality)
	appointmentRoute.GET("/:doctor_id", appointmentHandler.GetAppointmentsByDoctorId)
	appointmentRoute.GET("/speciality/:slug", appointmentHandler.GetAppointmentsBySpeciality)
	appointmentRoute.POST("/create-appointment", middleware.AuthMiddleware(constants.RoleAuthenticated), appointmentHandler.CreateAppointment)

}
