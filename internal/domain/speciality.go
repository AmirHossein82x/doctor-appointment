package domain

type Speciality struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"size:255;not null;unique"`
	Slug        string `gorm:"size:255;not null;unique"`
	Description string
}

// TableName specifies the table name for GORM
func (Speciality) TableName() string {
	return "speciality"
}
