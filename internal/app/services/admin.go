package services

import (
	"net/http"

	"github.com/AmirHossein82x/doctor-appointment/internal/app/dto"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/ports"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AdminService struct {
	adminRepository ports.AdminRepository
	log             *logrus.Logger
	smsService      ports.SmsService
}

func NewAdminService(adminRepository ports.AdminRepository, log *logrus.Logger, smsService ports.SmsService) *AdminService {
	return &AdminService{adminRepository: adminRepository, log: log, smsService: smsService}
}

// retrieve all users with pagination and search
// @Summary retrieve all users with pagination and search
// @Description retrieve all users with pagination and search
// @Tags admin
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param search query string false "Search query (name or phone starts with)"
// @Param role query string false "Search query (based on user_role)" Enums(admin, doctor, normal)
// @Security BearerAuth
// @Router /admin/get-all-users [get]
func (a *AdminService) GetAllUsers(c *gin.Context) {
	a.log.Info("Get all users with pagination and search")

	// Parse pagination parameters
	page, err := utils.GetQueryInt(c, "page", 1)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid page parameter")
		return
	}
	limit, err := utils.GetQueryInt(c, "limit", 10)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid limit parameter")
		return
	}

	// Parse search parameter
	search := c.Query("search")

	// Parse and validate role parameter
	role := c.Query("role")
	if role != "" && !utils.IsValidRole(role) {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid role parameter")
		return
	}

	// Fetch users with pagination and search
	users, err := a.adminRepository.GetAllUsers(page, limit, search, role)
	if err != nil {
		a.log.Error("Error getting all users: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, "Users retrieved successfully", users)
}

// create a new speciality
// @Summary Create a new speciality
// @Description Create a new speciality with a name, description, and auto-generated slug
// @Tags admin
// @Accept json
// @Produce json
// @Param speciality body dto.SpecialityCreateRequest true "Speciality details"
// @Security BearerAuth
// @Router /admin/create-speciality [post]
func (a *AdminService) CreateSpeciality(c *gin.Context) {
	a.log.Info("Create speciality")
	var req dto.SpecialityCreateRequest
	if err := c.BindJSON(&req); err != nil {
		a.log.Error("Invalid request")
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
		return
	}

	// Generate slug from name
	slug := utils.GenerateSlug(req.Name)

	speciality, err := a.adminRepository.CreateSpeciality(req.Name, slug, req.Description)
	if err != nil {
		a.log.Error("Error creating speciality: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, "Speciality created successfully", speciality)
}

// retrieve specialities with pagination and search
// @Summary Retrieve specialities with pagination and search
// @Description Retrieve specialities with pagination and search on the name of the speciality
// @Tags admin
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param search query string false "Search query (name starts with)"
// @Security BearerAuth
// @Router /admin/get-specialities [get]
func (a *AdminService) RetrieveSpeciality(c *gin.Context) {
	a.log.Info("Retrieve specialities with pagination and search")

	// Parse pagination parameters
	page, err := utils.GetQueryInt(c, "page", 1)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid page parameter")
		return
	}
	limit, err := utils.GetQueryInt(c, "limit", 10)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid limit parameter")
		return
	}

	// Parse search parameter
	search := c.Query("search")

	// Fetch specialities with pagination and search
	specialities, err := a.adminRepository.RetrieveSpeciality(page, limit, search)
	if err != nil {
		a.log.Error("Error retrieving specialities: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, "Specialities retrieved successfully", specialities)
}
