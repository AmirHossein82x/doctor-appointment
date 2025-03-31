package ports

import (
	"github.com/AmirHossein82x/doctor-appointment/internal/app/dto"
	"github.com/AmirHossein82x/doctor-appointment/internal/domain"
	"github.com/gin-gonic/gin"
)

type AdminService interface {
	GetAllUsers(*gin.Context)
	CreateSpeciality(*gin.Context)
	RetrieveSpeciality(*gin.Context)
}

type AdminRepository interface {
	GetAllUsers(int, int, string) ([]dto.UserRetrieveResponse, error)
	CreateSpeciality(string, string, string) (domain.Speciality, error)
	RetrieveSpeciality(int, int, string) ([]dto.SpecialityRetrieveResponse, error)
}
