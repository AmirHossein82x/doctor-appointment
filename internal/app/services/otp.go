package services

import (
	"fmt"

	"github.com/AmirHossein82x/doctor-appointment/internal/app/dto"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/validator"
	"github.com/AmirHossein82x/doctor-appointment/internal/repository"
	"github.com/AmirHossein82x/doctor-appointment/internal/sms_sender"
	"github.com/AmirHossein82x/doctor-appointment/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type OTPService struct {
	otpRepo    repository.OtpRepoInterface
	log        *logrus.Logger
	smsService sms_sender.SmsService
}

func NewOTPService(otpRepo repository.OtpRepoInterface, log *logrus.Logger, smsService sms_sender.SmsService) *OTPService {
	return &OTPService{otpRepo: otpRepo, log: log, smsService: smsService}
}

// Generate OTP
// @Summary Generate OTP
// @Description Generates an OTP for phone number verification
// @Tags OTP
// @Accept json
// @Produce json
// @Param request body dto.OTPRequest true "Phone Number"
// @Router /otp/generate [post]
func (o *OTPService) GenerateOTP(c *gin.Context) {
	o.log.Info("GenerateOTP")
	var req dto.OTPRequest
	if err := c.BindJSON(&req); err != nil {
		o.log.Error("Invalid request")
		utils.ErrorResponse(c, 400, "Invalid request")
		return

	}
	if !validator.ValidateIranianPhoneNumber(req.PhoneNumber) {
		o.log.Error("Invalid phone number")
		utils.ErrorResponse(c, 400, "Invalid phone number")
		return
	}
	otp, err := o.otpRepo.GenerateOTP(req.PhoneNumber)
	if err != nil {
		o.log.Error(err.Error())
		utils.ErrorResponse(c, 500, "can not generate otp")
		return
	}
	go o.smsService.SendSMS([]string{req.PhoneNumber}, fmt.Sprintf("Your OTP is: %d", otp))
	o.log.Info("OTP generated successfully")
	utils.SuccessResponse(c, "OTP generated successfully", nil)

}

// Verify OTP
// @Summary Verify OTP
// @Description Verifies the OTP entered by the user
// @Tags OTP
// @Accept json
// @Produce json
// @Param request body dto.VerifyOTPRequest true "Phone Number and OTP"
// @Router /otp/verify [post]
func (o *OTPService) VerifyOTP(c *gin.Context) {
	o.log.Info("VerifyOTP")
	var req dto.VerifyOTPRequest
	if err := c.BindJSON(&req); err != nil {
		o.log.Error("Invalid request")
		utils.ErrorResponse(c, 400, "Invalid request")
		return

	}
	if !validator.ValidateIranianPhoneNumber(req.PhoneNumber) {
		o.log.Error("Invalid phone number")
		utils.ErrorResponse(c, 400, "Invalid phone number")
		return
	}

	ok, err := o.otpRepo.VerifyOTP(req.PhoneNumber, req.OTP)
	if err != nil {
		o.log.Warn(err.Error())
		utils.ErrorResponse(c, 500, "can not verify otp")
		return
	}
	if !ok {
		o.log.Warn("Invalid OTP")
		utils.ErrorResponse(c, 400, "Invalid OTP")
		return
	}
	o.log.Info("OTP verified successfully")
	token, err := o.otpRepo.GenerateVerifyOtpToken(req.PhoneNumber)
	if err != nil {
		o.log.Error("can not generate verified token")
		utils.ErrorResponse(c, 500, "can not generate verified token")
		return
	}
	utils.SuccessResponse(c, "OTP verified successfully", gin.H{"verified_token": token})
}
