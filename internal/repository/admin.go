package repository

import (
	"github.com/AmirHossein82x/doctor-appointment/internal/app/dto"
	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{DB: infrastructure.DB}
}

func (a *AdminRepository) GetAllUsers(page int, limit int, search string) ([]dto.UserRetrieveResponse, error) {
	var users []dto.UserRetrieveResponse
	offset := (page - 1) * limit

	query := a.DB.Table("users").
		Select("id, phone, name, role")

	// Add search condition if search parameter is provided
	if search != "" {
		query = query.Where("phone LIKE ? OR name LIKE ?", search+"%", search+"%")
	}

	// Apply pagination after filtering
	query = query.Offset(offset).Limit(limit)

	err := query.Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
