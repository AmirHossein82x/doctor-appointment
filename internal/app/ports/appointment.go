package ports

import (
	"github.com/AmirHossein82x/doctor-appointment/internal/app/dto"
	"github.com/gin-gonic/gin"
)


type AppointmentService interface {
	GetDoctorProfiles(*gin.Context)
	RetrieveSpeciality(*gin.Context)
}


type AppointmentRepository interface {
	GetDoctorProfiles(int, int, string) ([]map[string]interface{}, error)
	RetrieveSpeciality(int, int, string) (*[]dto.SpecialityRetrieveResponse, error)
}