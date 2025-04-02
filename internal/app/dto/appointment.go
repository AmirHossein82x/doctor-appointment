package dto

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID        uuid.UUID `json:"id"`
	Date      time.Time `json:"date" binding:"required"`
	StartTime string    `json:"start_time" binding:"required"` // Expecting HH:MM:SS format
	EndTime   string    `json:"end_time" binding:"required"`   // Expecting HH:MM:SS format
	Status    string    `json:"status"`
}
