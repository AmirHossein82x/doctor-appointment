package utils

import (
	"fmt"
	"time"

	"github.com/AmirHossein82x/doctor-appointment/internal/app/constants"
	"github.com/AmirHossein82x/doctor-appointment/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var cfg = config.LoadConfig()
var mySigningKey = []byte(cfg.SECRET_KEY)

// Create a struct to represent the claims
type Claims struct {
	ID        uuid.UUID           `json:"id"`
	Name      string              `json:"name"`
	Role      string              `json:"role"`
	TokenType constants.TokenType `json:"token_type"`
	jwt.RegisteredClaims
}

func GenerateToken(ID uuid.UUID, Name string, Role string, TokenType constants.TokenType) (string, error) {
	// Create a new token
	var token *jwt.Token
	if TokenType == constants.TokenTypeAccess {
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
			ID:        ID,
			Name:      Name,
			Role:      Role,
			TokenType: constants.TokenTypeAccess,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(constants.AccessTokenLifetime)),
			},
		})
	} else {
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
			ID:        ID,
			Name:      Name,
			Role:      Role,
			TokenType: constants.TokenTypeRefresh,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(constants.RefreshTokenLifetime)),
			},
		})
	}

	// Sign the token with our secret
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err // Return an empty string and the error
	}

	// Return the generated token and no error
	return tokenString, nil
}

func VerifyToken(Token string, TokenType constants.TokenType) (*Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(Token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
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

	if claims.TokenType != TokenType {
		return nil, fmt.Errorf("invalid token type")
	}

	return claims, nil
}
