package constants

type RoleType string

const (
	AdminRole  RoleType = "admin"
	NormalRole RoleType = "normal"
	DoctorRole RoleType = "doctor"
	
	BearerPrefix = "Bearer "
	
)
