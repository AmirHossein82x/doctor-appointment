package repository

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type OtpRepository struct {
	redisClient *redis.Client
}

func NewOtpRepo(redisClient *redis.Client) *OtpRepository {
	return &OtpRepository{redisClient: redisClient}
}
func (o *OtpRepository) GenerateOTP(phoneNumber string) (int, error) {
	// Create a context with timeout to prevent long-running Redis operations
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Generate a random OTP in the range [10000, 99999]
	otp := rand.Intn(90000) + 10000

	// Store OTP in Redis with expiration time of 5 minutes
	err := o.redisClient.Set(ctx, phoneNumber, otp, 5*time.Minute).Err()
	if err != nil {
		return 0, fmt.Errorf("failed to store OTP in Redis: %w", err)
	}

	return otp, nil
}

func (o *OtpRepository) VerifyOTP(PhoneNumber string, otp string) (bool, error) {
	// Create a context with timeout to prevent long-running Redis operations
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Retrieve OTP from Redis
	storedOtp, err := o.redisClient.Get(ctx, PhoneNumber).Result()
	if err != nil {
		return false, fmt.Errorf("failed to retrieve OTP from Redis: %w", err)
	}

	// Compare the OTPs
	if storedOtp == otp {
		return true, nil
	}

	return false, nil
}

func (o *OtpRepository) GenerateVerifyOtpToken(PhoneNumber string) (string, error) {
	// Create a context with timeout to prevent long-running Redis operations
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	verifiedToken := uuid.NewString()
	err := o.redisClient.Set(ctx, verifiedToken, PhoneNumber, 5*time.Minute).Err()
	if err != nil {
		return "", fmt.Errorf("failed to store verified token in Redis: %w", err)
	}
	return verifiedToken, nil

}
