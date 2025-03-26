package utils

import (
	"fmt"
	"time"

	"github.com/AmirHossein82x/doctor-appointment/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var cfg = config.LoadConfig()
var mySigningKey = []byte(cfg.SECRET_KEY)

// Create a struct to represent the claims
type Claims struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	TokenType string    `json:"token_type"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(ID uuid.UUID, Name string, Role string) (string, error) {
	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		ID:        ID,
		Name:      Name,
		Role:      Role,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)), // Token expires in 30 minutes
		},
	})

	// Sign the token with our secret
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err // Return an empty string and the error
	}

	// Return the generated token and no error
	return tokenString, nil
}

func GenerateRefreshToken(ID uuid.UUID, Name string, Role string) (string, error) {
	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		ID:        ID,
		Name:      Name,
		Role:      Role,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 1 day
		},
	})

	// Sign the token with our secret
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err // Return an empty string and the error
	}

	// Return the generated token and no error
	return tokenString, nil
}

func VerifyAccessToken(accessToken string) (*Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Extract the claims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}

	if claims.TokenType != "access" {
		return nil, fmt.Errorf("invalid token type")
	}

	return claims, nil
}



func VerifyRefreshToken(accessToken string) (*Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Extract the claims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}

	if claims.TokenType != "refresh" {
		return nil, fmt.Errorf("invalid token type")
	}

	return claims, nil
}
