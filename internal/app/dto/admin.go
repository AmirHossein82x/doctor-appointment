package dto

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
