package ports

import "github.com/gin-gonic/gin"


type AppointmentService interface {
	GetDoctorProfiles(*gin.Context)
}


type AppointmentRepository interface {
	GetDoctorProfiles(int, int) ([]map[string]interface{}, error)
}