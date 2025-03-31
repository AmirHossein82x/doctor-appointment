package ports

import (
	"github.com/AmirHossein82x/doctor-appointment/internal/app/dto"
	"github.com/gin-gonic/gin"
)


type AdminService interface {
	GetAllUsers(*gin.Context) 
	
}

type AdminRepository interface {
	GetAllUsers(int, int, string) ([]dto.UserRetrieveResponse, error)
}
