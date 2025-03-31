package utils

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetQueryInt(c *gin.Context, key string, defaultValue int) (int, error) {
	valueStr := c.Query(key)
	if valueStr == "" {
		return defaultValue, nil
	}
	return strconv.Atoi(valueStr)
}

// GenerateSlug generates a slug from a given string, supporting Persian letters
func GenerateSlug(input string) string {
	// Convert to lowercase
	input = strings.ToLower(input)

	// Replace spaces with hyphens
	input = strings.ReplaceAll(input, " ", "-")

	// Remove invalid characters (allow Persian letters, English letters, numbers, and hyphens)
	// reg := regexp.MustCompile(`[^a-z0-9\u0600-\u06FF-]`)
	// input = reg.ReplaceAllString(input, "")

	return input
}

var allowedRoles = []string{"admin", "doctor", "normal"}

// IsValidRole checks if the given role is valid
func IsValidRole(role string) bool {
	for _, allowedRole := range allowedRoles {
		if role == allowedRole {
			return true
		}
	}
	return false
}
