package services

import (
	"net/http"

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

// retrieve all users with pagination
// @Summary retrieve all users with pagination
// @Description retrieve all users with pagination
// @Tags admin
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Security BearerAuth
// @Router /admin/get-all-users [get]
func (a *AdminService) GetAllUsers(c *gin.Context) {
	a.log.Info("Get all users with pagination")

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

	// Fetch users with pagination
	users, err := a.adminRepository.GetAllUsers(page, limit)
	if err != nil {
		a.log.Error("Error getting all users: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, "Users retrieved successfully", users)
}
