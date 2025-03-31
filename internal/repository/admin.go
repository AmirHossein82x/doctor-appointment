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

func (a *AdminRepository) GetAllUsers(page int, limit int) ([]dto.UserResponse, error) {
	var users []dto.UserResponse
	offset := (page - 1) * limit
	err := a.DB.Table("users").
		Select("id, phone, name, role").
		Offset(offset).
		Limit(limit).
		Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
