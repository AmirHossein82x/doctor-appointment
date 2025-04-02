package ports

import (
	"time"

	"github.com/AmirHossein82x/doctor-appointment/internal/app/dto"
	"github.com/AmirHossein82x/doctor-appointment/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AppointmentService interface {
	GetDoctorProfiles(*gin.Context)
	RetrieveSpeciality(*gin.Context)
	GetAppointmentsByDoctorId(*gin.Context)
	GetAppointmentsBySpeciality(*gin.Context)
}

type AppointmentRepository interface {
	GetDoctorProfiles(int, int, string) ([]map[string]interface{}, error)
	RetrieveSpeciality(int, int, string) (*[]dto.SpecialityRetrieveResponse, error)
	GetAppointmentsByDoctorId(uuid.UUID, time.Time, string, int, int) (*[]domain.DoctorAppointment, error)
	GetAppointmentsBySpeciality(string, time.Time, string, int, int) ([]map[string]interface{}, error)
}
