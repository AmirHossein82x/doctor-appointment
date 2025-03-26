package domain

import (
	"time"
	"github.com/google/uuid"
)

// User represents the structure of the users table
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"` // UUID as primary key
	Name      string    `gorm:"type:varchar(100);not null"`
	Password  string    `gorm:"type:text;not null"`
	Phone     string    `gorm:"type:varchar(20);unique"`
	Role      string    `gorm:"type:user_role;not null;default:'normal'"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "users"
}