package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/AmirHossein82x/doctor-appointment/internal/domain"
	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthRepository struct {
	redisClient *redis.Client
	DB          *gorm.DB
}

func NewAuthRepository(redisClient *redis.Client) *AuthRepository {
	return &AuthRepository{redisClient: redisClient, DB: infrastructure.DB}
}

func (u *AuthRepository) GetPhoneNumberFromToken(token string) (string, error) {
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

func (u *AuthRepository) Register(user *domain.User) error {
	err := u.DB.Create(user).Error
	return err

}

func (u *AuthRepository) GetByPhoneNumber(phone string) (domain.User, error) {
	var user domain.User
	err := u.DB.Where("phone = ?", phone).First(&user).Error
	return user, err
}

func (u *AuthRepository) UpdatePassword(id uuid.UUID, Password string) error {
	err := u.DB.Model(&domain.User{}).Where("id = ?", id).Update("password", Password).Error
	return err
}

func (u *AuthRepository) SaveEncryptionKeyToRedis(encryptionKey string) error {
	// Create a context with timeout to prevent long-running Redis operations
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Store OTP in Redis with expiration time of 5 minutes
	err := u.redisClient.Set(ctx, encryptionKey, true, 5*time.Minute).Err()
	return err
}

func (u *AuthRepository) ExistsEncryptionKey(encryptionKey string) bool {
	// Create a context with timeout to prevent long-running Redis operations
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := u.redisClient.Get(ctx, encryptionKey).Result()
	return err == nil
}

func (u *AuthRepository) DeleteEncryptionKey(encryptionKey string) error {
	// Create a context with timeout to prevent long-running Redis operations
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Use Redis Del command to delete the key
	_, err := u.redisClient.Del(ctx, encryptionKey).Result()
	return err
}
