package repository

import (
	"github.com/AmirHossein82x/doctor-appointment/internal/app/dto"
	"github.com/AmirHossein82x/doctor-appointment/internal/domain"
	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{DB: infrastructure.DB}
}

func (a *AdminRepository) GetAllUsers(page int, limit int, search string, role string) ([]dto.UserRetrieveResponse, error) {
	var users []dto.UserRetrieveResponse
	offset := (page - 1) * limit

	query := a.DB.Table("users").
		Select("id, phone, name, role")

	// Add search condition if search parameter is provided
	if search != "" {
		query = query.Where("phone LIKE ? OR name LIKE ?", search+"%", search+"%")
	}
	if role != "" {
		query = query.Where("role = ?", role)
	}

	// Apply pagination after filtering
	query = query.Offset(offset).Limit(limit)

	err := query.Scan(&users).Error
	return users, err
}

func (a *AdminRepository) CreateSpeciality(name, slug, description string) (domain.Speciality, error) {
	var speciality domain.Speciality = domain.Speciality{
		Name:        name,
		Slug:        slug,
		Description: description,
	}
	err := a.DB.Create(&speciality).Error
	return speciality, err

}

func (a *AdminRepository) RetrieveSpeciality(page int, limit int, search string) ([]dto.SpecialityRetrieveResponse, error) {
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
	return specialities, err
}

func (a AdminRepository) UpdateRole(userID uuid.UUID) error {
	var user domain.User
	err := a.DB.Model(&user).Where("id = ?", userID).Update("role", "doctor").Error
	return err
}

func (a *AdminRepository) CreateDoctorProfile(req dto.DoctorProfileCreateRequest) (domain.DoctorProfile, error) {
	var doctorProfile domain.DoctorProfile = domain.DoctorProfile{
		SpecialityID:    req.SpecialityID,
		Bio:             req.Bio,
		ExperienceYears: req.ExperienceYears,
	}
	err := a.DB.Create(&doctorProfile).Error
	return doctorProfile, err

}

func (a *AdminRepository) CreateDoctorProfileWithTransaction(req dto.DoctorProfileCreateRequest) (domain.DoctorProfile, error) {
	var doctorProfile domain.DoctorProfile

	// Start a transaction
	tx := a.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create the doctor profile
	doctorProfile = domain.DoctorProfile{
		ID:              req.UserID,
		SpecialityID:    req.SpecialityID,
		Bio:             req.Bio,
		ExperienceYears: req.ExperienceYears,
	}
	if err := tx.Create(&doctorProfile).Error; err != nil {
		tx.Rollback()
		return doctorProfile, err
	}

	// Update the user's role to "doctor"
	if err := tx.Model(&domain.User{}).Where("id = ?", req.UserID).Update("role", "doctor").Error; err != nil {
		tx.Rollback()
		return doctorProfile, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return doctorProfile, err
	}

	return doctorProfile, nil
}
