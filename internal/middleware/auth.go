package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/AmirHossein82x/doctor-appointment/internal/app/constants"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/utils"
	"github.com/AmirHossein82x/doctor-appointment/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey []byte

// Initialize the signing key using LoadConfig
func init() {
	cfg := config.LoadConfig()
	mySigningKey = []byte(cfg.SECRET_KEY)
}

// AuthMiddleware checks if the user is authenticated
func AuthMiddleware(Role constants.RoleType) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			handleUnauthorized(c, "Authorization header is required")
			return
		}

		// Extract the token from the header
		if !strings.HasPrefix(authHeader, constants.BearerPrefix) {
			handleUnauthorized(c, "Invalid token format")
			return
		}
		tokenString := strings.TrimPrefix(authHeader, constants.BearerPrefix)

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return mySigningKey, nil
		})

		if err != nil || !token.Valid {
			handleUnauthorized(c, "Invalid token")
			return
		}

		// Validate token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			handleUnauthorized(c, "Invalid token claims")
			return
		}

		tokenType, ok := claims["token_type"].(string)
		if !ok || tokenType != string(constants.TokenTypeAccess) {
			handleUnauthorized(c, "Invalid token type")
			return
		}

		role, ok := claims["role"].(string)
		if !ok || role != string(Role) {
			handleUnauthorized(c, "you do not have permission to access this resource")
			return
		}

		// Add the username to the context for use in subsequent handlers
		c.Set("id", claims["id"])
		c.Set("name", claims["name"])
		c.Set("role", claims["role"])

		c.Next()
	}
}

// handleUnauthorized centralizes error responses for unauthorized access
func handleUnauthorized(c *gin.Context, message string) {
	utils.ErrorResponse(c, http.StatusUnauthorized, message)
	c.Abort()
}
