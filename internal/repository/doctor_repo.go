package repository

import (
	"time"

	"github.com/AmirHossein82x/doctor-appointment/internal/domain"
	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DoctorRepository struct {
	DB *gorm.DB
}

func NewDoctorRepository() *DoctorRepository {
	return &DoctorRepository{DB: infrastructure.DB}
}

func (d *DoctorRepository) CreateAppointment(date, startTime, endTime time.Time, doctorId uuid.UUID) (*domain.DoctorAppointment, error) {
	appointment := domain.DoctorAppointment{
		Date:            date,
		StartTime:       startTime,
		EndTime:         endTime,
		DoctorProfileID: doctorId,
	}
	err := d.DB.Create(&appointment).Error
	return &appointment, err

}
func (d *DoctorRepository) IsAppointmentAvailable(date, startTime, endTime time.Time, doctorId uuid.UUID) (bool, error) {
	var count int64
	err := d.DB.Model(&domain.DoctorAppointment{}).
		Where("doctor_profile_id = ? AND date = ? AND start_time < ? AND end_time > ?",
			doctorId, date, endTime, startTime).
		Count(&count).Error

	return count == 0, err
}
