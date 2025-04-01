package domain

import (
	"time"

	"github.com/google/uuid"
)

type DoctorProfile struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	SpecialityID    int       `gorm:"not null"`
	Bio             string
	ExperienceYears int       `gorm:"check:experience_years >= 0"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}

// TableName specifies the table name for GORM
func (DoctorProfile) TableName() string {
	return "doctor_profile"
}

type DoctorAppointment struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"` // UUID as primary key
	DoctorProfileID uuid.UUID `gorm:"type:uuid;not null"`
	Date            time.Time `gorm:"type:date;not null"`
	StartTime       time.Time `gorm:"type:time;not null"`
	EndTime         time.Time `gorm:"type:time;not null"`
	Status          string    `gorm:"type:varchar(50);not null;default:'available';check:status IN ('available', 'booked')"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

// TableName specifies the table name for GORM
func (DoctorAppointment) TableName() string {
	return "doctor_appointment"
}
