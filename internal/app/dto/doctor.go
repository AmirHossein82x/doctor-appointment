package dto


type AppointmentCreateRequest struct {
	Date      string `json:"date" binding:"required"`       // Expecting YYYY-MM-DD format
	StartTime string `json:"start_time" binding:"required"` // Expecting HH:MM:SS format
	EndTime   string `json:"end_time" binding:"required"`   // Expecting HH:MM:SS format
}