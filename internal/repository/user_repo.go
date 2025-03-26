package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/AmirHossein82x/doctor-appointment/internal/domain"
	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	Register(*domain.User) error
	GetPhoneNumberFromToken(string) (string, error)
}

type UserRepository struct {
	redisClient *redis.Client
	DB          *gorm.DB
}

func NewUserRepository(redisClient *redis.Client) *UserRepository {
	return &UserRepository{redisClient: redisClient, DB: infrastructure.DB}
}

func (u *UserRepository) GetPhoneNumberFromToken(token string) (string, error) {
	// Create a context with timeout to prevent long-running Redis operations
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Retrieve OTP from Redis
	phoneNumber, err := u.redisClient.Get(ctx, token).Result()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve token from Redis: %w", err)
	}

	return phoneNumber, nil
}

func (u *UserRepository) Register(user *domain.User) error{
	err := u.DB.Create(user).Error
	return err

}
