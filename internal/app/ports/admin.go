package ports

import (
	"github.com/AmirHossein82x/doctor-appointment/internal/app/dto"
	"github.com/AmirHossein82x/doctor-appointment/internal/domain"
	"github.com/gin-gonic/gin"
)

type AdminService interface {
	GetAllUsers(*gin.Context)
	CreateSpeciality(*gin.Context)
	CreateDoctorProfile(*gin.Context)
}

type AdminRepository interface {
	GetAllUsers(page int, limit int, search string, role string) ([]dto.UserRetrieveResponse, error)
	CreateSpeciality(string, string, string) (domain.Speciality, error)
	CreateDoctorProfileWithTransaction(*dto.DoctorProfileCreateRequest) (domain.DoctorProfile, error)
}
