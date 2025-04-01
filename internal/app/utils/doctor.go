package utils

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// parseAppointmentTimes parses and combines date and time strings into full time.Time objects.
func ParseAppointmentTimes(dateStr, startStr, endStr string) (time.Time, time.Time, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	startTime, err := time.Parse("15:04:05", startStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	endTime, err := time.Parse("15:04:05", endStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	// Combine date with times
	start := time.Date(date.Year(), date.Month(), date.Day(), startTime.Hour(), startTime.Minute(), startTime.Second(), 0, date.Location())
	end := time.Date(date.Year(), date.Month(), date.Day(), endTime.Hour(), endTime.Minute(), endTime.Second(), 0, date.Location())

	if !start.Before(end) {
		return time.Time{}, time.Time{}, errors.New("start time must be before end time")
	}

	return start, end, nil
}

// getDoctorID retrieves and parses the doctor ID from the Gin context.
func GetDoctorID(c *gin.Context) (uuid.UUID, error) {
	id, exists := c.Get("id")
	if !exists {
		return uuid.Nil, errors.New("user ID not found in context")
	}

	idStr, ok := id.(string)
	if !ok {
		return uuid.Nil, errors.New("invalid user ID format")
	}

	return uuid.Parse(idStr)
}