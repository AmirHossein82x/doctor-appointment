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
