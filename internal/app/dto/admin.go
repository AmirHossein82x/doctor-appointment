package dto

import "github.com/google/uuid"

type UserRetrieveResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

type SpecialityCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type SpecialityCreateResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description"`
}

type SpecialityRetrieveResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

type DoctorProfileCreateRequest struct {
	UserID          uuid.UUID `json:"user_id" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	SpecialityID    int       `json:"speciality_id" binding:"required"`
	Bio             string    `json:"bio"`
	ExperienceYears int       `json:"experience_years" binding:"required"`
}
