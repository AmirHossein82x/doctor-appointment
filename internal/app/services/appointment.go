package services

import (
	"net/http"

	"github.com/AmirHossein82x/doctor-appointment/internal/app/ports"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AppointmentService struct {
	appointmentRepo ports.AppointmentRepository
	log             *logrus.Logger
	smsService      ports.SmsService
}

func NewAppointmentService(appointmentRepo ports.AppointmentRepository, log *logrus.Logger, smsService ports.SmsService) *AppointmentService {
	return &AppointmentService{appointmentRepo: appointmentRepo, log: log, smsService: smsService}
}

// retrieve all doctor profiles joined with user table
// @Summary Retrieve all doctor profiles
// @Description Retrieve all doctor profiles joined with user table
// @Tags appointment
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param search query string false "Search query (slug of speciality)"
// @Router /appointment/get-doctor-profiles [get]
func (a *AppointmentService) GetDoctorProfiles(c *gin.Context) {
	a.log.Info("Get Doctor Profiles called")

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
	search := c.Query("search")

	// Fetch doctor profiles
	doctorProfiles, err := a.appointmentRepo.GetDoctorProfiles(page, limit, search)
	if err != nil {
		a.log.Error("Error retrieving doctor profiles: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, "Doctor profiles retrieved successfully", doctorProfiles)
}


// retrieve specialities with pagination and search
// @Summary Retrieve specialities with pagination and search
// @Description Retrieve specialities with pagination and search on the name of the speciality
// @Tags appointment
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param search query string false "Search query (name starts with)"
// @Router /appointment/get-specialities [get]
func (a *AppointmentService) RetrieveSpeciality(c *gin.Context) {
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
	specialities, err := a.appointmentRepo.RetrieveSpeciality(page, limit, search)
	if err != nil {
		a.log.Error("Error retrieving specialities: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, "Specialities retrieved successfully", specialities)
}