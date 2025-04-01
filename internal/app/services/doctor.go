package services

import (
	"net/http"

	"github.com/AmirHossein82x/doctor-appointment/internal/app/dto"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/ports"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DoctorService struct {
	doctorRepository ports.DoctorRepository
	log              *logrus.Logger
	smsService       ports.SmsService
}

func NewDoctorService(doctorRepository ports.DoctorRepository, log *logrus.Logger, smsService ports.SmsService) *DoctorService {
	return &DoctorService{doctorRepository: doctorRepository, log: log, smsService: smsService}
}

// CreateAppointment handles the creation of a doctor appointment.
// @Summary create doctor appointment
// @Description create doctor appointment
// @Tags doctor
// @Accept json
// @Produce json
// @Param doctorProfile body dto.AppointmentCreateRequest true "Doctor appointment details"
// @Security BearerAuth
// @Router /doctor/create-appointment [post]
func (d *DoctorService) CreateAppointment(c *gin.Context) {
	var req dto.AppointmentCreateRequest
	if err := c.BindJSON(&req); err != nil {
		d.log.Error("bad request")
		utils.ErrorResponse(c, http.StatusBadRequest, "failed to bin json")
		return
	}

	// Parse and validate appointment times
	start, end, err := utils.ParseAppointmentTimes(req.Date, req.StartTime, req.EndTime)
	if err != nil {
		d.log.Error("time parsing failed")
		utils.ErrorResponse(c, http.StatusBadRequest, "time parsing failed")
		return
	}

	// Get and parse Doctor ID from context
	doctorID, err := utils.GetDoctorID(c)
	if err != nil {

		d.log.Error("failed to get doctor id")
		utils.ErrorResponse(c, http.StatusBadRequest, "failed to get doctor id")
		return
	}

	// Check availability
	available, err := d.doctorRepository.IsAppointmentAvailable(start, start, end, doctorID)
	if err != nil {
		d.log.Error("some thing went wrong")
		utils.ErrorResponse(c, http.StatusBadRequest, "some thing went wrong")
		return
	}
	if !available {
		d.log.Error("appointment time is already booked")
		utils.ErrorResponse(c, http.StatusConflict, "appointment time is already booked")
		return
	}

	// Create appointment
	appointment, err := d.doctorRepository.CreateAppointment(start, start, end, doctorID)
	if err != nil {
		d.log.Error("failed to create appointment")
		utils.ErrorResponse(c, http.StatusBadRequest, "failed to create appointment")
		return
	}

	utils.SuccessResponse(c, "appointment created successfully", appointment)
}

// get all available appointments
// @Summary get all available appointments
// @Description get all available appointments
// @Tags doctor
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Security BearerAuth
// @Router /doctor/available-appointments [get]
func (d *DoctorService) GetAvailableAppointments(c *gin.Context) {
	// Get and parse Doctor ID from context
	doctorID, err := utils.GetDoctorID(c)
	if err != nil {
		d.log.Error("failed to get doctor id")
		utils.ErrorResponse(c, http.StatusBadRequest, "failed to get doctor id")
		return
	}
	

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
	appointments, err := d.doctorRepository.GetAvailableAppointments(doctorID, page, limit)
	if err != nil {
		d.log.Error("internal server error")
		utils.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}
	utils.SuccessResponse(c, "retrieved suucessfully", appointments)
}
