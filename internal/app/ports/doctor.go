package ports

import (
	"time"

	"github.com/AmirHossein82x/doctor-appointment/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DoctorService interface {
	CreateAppointment(*gin.Context)
	GetAvailableAppointments(*gin.Context)
	GetBookedAppointments(*gin.Context)
}

type DoctorRepository interface {
	CreateAppointment(time.Time, time.Time, time.Time, uuid.UUID) (domain.DoctorAppointment, error)
	IsAppointmentAvailable(time.Time, time.Time, time.Time, uuid.UUID) (bool, error)
	GetAvailableAppointments(uuid.UUID, int, int) ([]domain.DoctorAppointment, error)
	GetBookedAppointments(uuid.UUID, string, time.Time, int, int) ([]map[string]interface{}, error)
}
