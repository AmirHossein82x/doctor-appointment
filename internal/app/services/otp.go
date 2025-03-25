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
	utils.SuccessResponse(c, "OTP generated successfully", nil)

}

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
	utils.SuccessResponse(c, "OTP verified successfully", nil)
}
