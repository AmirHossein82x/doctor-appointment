package repository

import (
	"time"

	"github.com/AmirHossein82x/doctor-appointment/internal/app/dto"
	"github.com/AmirHossein82x/doctor-appointment/internal/domain"
	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppointmentRepository struct {
	DB *gorm.DB
}

func NewAppointmentRepository() *AppointmentRepository {
	return &AppointmentRepository{DB: infrastructure.DB}
}

func (a *AppointmentRepository) GetDoctorProfiles(page int, limit int, search string) ([]map[string]interface{}, error) {
	var doctorProfiles []map[string]interface{}
	offset := (page - 1) * limit

	query := a.DB.Table("doctor_profile").
		Select("doctor_profile.id, doctor_profile.bio, doctor_profile.experience_years, users.name, users.phone, users.role, speciality.name as speciality_name").
		Joins("JOIN users ON doctor_profile.id = users.id").
		Joins("JOIN speciality ON doctor_profile.speciality_id = speciality.id")
	if search != "" {
		query = query.Where("speciality.slug = ?", search)
	}
	query = query.Offset(offset).Limit(limit)
	err := query.Scan(&doctorProfiles).Error
	return doctorProfiles, err
}

func (a *AppointmentRepository) RetrieveSpeciality(page int, limit int, search string) (*[]dto.SpecialityRetrieveResponse, error) {
	var specialities []dto.SpecialityRetrieveResponse
	offset := (page - 1) * limit

	query := a.DB.Table("speciality").
		Select("id, name, slug, description")

	// Add search condition if search parameter is provided
	if search != "" {
		query = query.Where("name LIKE ?", search+"%")
	}

	// Apply pagination after filtering
	query = query.Offset(offset).Limit(limit)

	err := query.Scan(&specialities).Error
	return &specialities, err
}

func (a *AppointmentRepository) GetAppointmentsByDoctorId(doctorId uuid.UUID, date time.Time, page, limit int) (*[]domain.DoctorAppointment, error) {
	var appointments []domain.DoctorAppointment
	offset := (page - 1) * limit
	now := time.Now()

	query := a.DB.Table("doctor_appointment").
		Where("doctor_profile_id = ? AND date >= ?", doctorId, now)
	if !date.IsZero() {
		query = query.Where("date = ?", date)
	}
	err := query.Offset(offset).
		Limit(limit).
		Find(&appointments).Error

	return &appointments, err
}
