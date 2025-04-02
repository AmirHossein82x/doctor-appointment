package ports

import (
	"github.com/AmirHossein82x/doctor-appointment/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthService interface {
	Register(*gin.Context)
	Login(*gin.Context)
	VerifyAccessToken(*gin.Context)
	GetAccessTokenByRefreshToken(*gin.Context)
	ForgetPassword(*gin.Context)
	ResetPassword(*gin.Context)
}

type AuthRepository interface {
	Register(*domain.User) error
	GetPhoneNumberFromToken(string) (string, error)
	GetByPhoneNumber(string) (domain.User, error)
	UpdatePassword(uuid.UUID, string) error
	SaveEncryptionKeyToRedis(string) error
	ExistsEncryptionKey(string) bool
	DeleteEncryptionKey(string) error
}
