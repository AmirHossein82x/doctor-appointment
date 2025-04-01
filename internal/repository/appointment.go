package repository

import (
	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
	"gorm.io/gorm"
)

type AppointmentRepository struct {
	DB *gorm.DB
}

func NewAppointmentRepository() *AppointmentRepository {
	return &AppointmentRepository{DB: infrastructure.DB}
}

func (a *AppointmentRepository) GetDoctorProfiles(page int, limit int) ([]map[string]interface{}, error) {
	var doctorProfiles []map[string]interface{}
	offset := (page - 1) * limit

	// Query to join doctor_profile and users table
	err := a.DB.Table("doctor_profile").
		Select("doctor_profile.id, doctor_profile.bio, doctor_profile.experience_years, users.name, users.phone, users.role, speciality.name as speciality_name").
		Joins("JOIN users ON doctor_profile.id = users.id").
		Joins("JOIN speciality ON doctor_profile.speciality_id = speciality.id").
		Offset(offset).
		Limit(limit).
		Scan(&doctorProfiles).Error

	if err != nil {
		return nil, err
	}
	return doctorProfiles, nil
}
