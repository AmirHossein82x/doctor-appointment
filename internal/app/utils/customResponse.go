package utils


import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// Response structure for consistent API responses
type Response struct {
    Msg  string      `json:"msg"`
    Data interface{} `json:"data,omitempty"` // Omits 'data' if nil
}

// SuccessResponse sends a successful JSON response
func SuccessResponse(c *gin.Context, msg string, data interface{}) {
    c.JSON(http.StatusOK, Response{
        Msg:  msg,
        Data: data,
    })
}

// ErrorResponse sends an error JSON response
func ErrorResponse(c *gin.Context, statusCode int, msg string) {
    c.JSON(statusCode, Response{
        Msg: msg,
    })
}
