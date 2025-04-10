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
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"` // UUID as primary key
	DoctorProfileID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Date            time.Time `gorm:"type:date;not null" json:"date"`
	StartTime       string    `gorm:"type:time;not null" json:"start_time"`
	EndTime         string    `gorm:"type:time;not null" json:"end_time"`
	Status          string    `gorm:"type:varchar(50);not null;default:'available';check:status IN ('available', 'booked')" json:"status"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName specifies the table name for GORM
func (DoctorAppointment) TableName() string {
	return "doctor_appointment"
}

type UserAppointment struct {
	ID                  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"-"` // UUID as primary key
	UserId              uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	DoctorAppointmentId uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Status              string    `gorm:"type:varchar(50);not null;default:'available';check:status IN ('available', 'booked')" json:"status"`
	CreatedAt           time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName specifies the table name for GORM
func (UserAppointment) TableName() string {
	return "user_appointment"
}
