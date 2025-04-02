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

func (a *AppointmentRepository) GetAppointmentsByDoctorId(doctorId uuid.UUID, date time.Time, status string, page, limit int) (*[]domain.DoctorAppointment, error) {
	var appointments []domain.DoctorAppointment
	offset := (page - 1) * limit
	now := time.Now()

	query := a.DB.Table("doctor_appointment").
		Where("doctor_profile_id = ? AND date >= ?", doctorId, now)

	// Add optional date filter
	if !date.IsZero() {
		query = query.Where("date = ?", date)
	}

	// Add optional status filter
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Offset(offset).
		Limit(limit).
		Find(&appointments).Error

	return &appointments, err
}

func (a *AppointmentRepository) GetAppointmentsBySpeciality(slug string, date time.Time, status string, page, limit int) ([]map[string]interface{}, error) {
	var appointments []map[string]interface{}
	offset := (page - 1) * limit

	// Build the base query
	query := a.DB.Table("doctor_appointment").
		Select(`
			users.name AS doctor_name,
			speciality.name AS speciality_name,
			doctor_appointment.date,
			doctor_appointment.start_time,
			doctor_appointment.end_time,
			doctor_appointment.status
		`).
		Joins("JOIN doctor_profile ON doctor_profile.id = doctor_appointment.doctor_profile_id").
		Joins("JOIN speciality ON speciality.id = doctor_profile.speciality_id").
		Joins("JOIN users ON doctor_profile.id = users.id").
		Where("speciality.slug = ?", slug).
		Where("doctor_appointment.date >= ?", time.Now())

	// Add optional date filter
	if !date.IsZero() {
		query = query.Where("doctor_appointment.date = ?", date)
	}

	// Add optional status filter
	if status != "" {
		query = query.Where("doctor_appointment.status = ?", status)
	}

	// Apply pagination
	query = query.Offset(offset).Limit(limit)

	// Execute the query
	err := query.Scan(&appointments).Error
	if err != nil {
		return nil, err
	}

	return appointments, nil
}
