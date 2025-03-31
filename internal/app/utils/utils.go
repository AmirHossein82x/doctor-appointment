package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetQueryInt(c *gin.Context, key string, defaultValue int) (int, error) {
	valueStr := c.Query(key)
	if valueStr == "" {
		return defaultValue, nil
	}
	return strconv.Atoi(valueStr)
}
