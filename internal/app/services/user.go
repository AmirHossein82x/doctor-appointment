package services

import (
	"time"

	"github.com/AmirHossein82x/doctor-appointment/internal/app/dto"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/utils"
	"github.com/AmirHossein82x/doctor-appointment/internal/domain"
	"github.com/AmirHossein82x/doctor-appointment/internal/repository"
	"github.com/AmirHossein82x/doctor-appointment/internal/sms_sender"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	userRepo   repository.UserRepoInterface
	log        *logrus.Logger
	smsService sms_sender.SmsService
}

func NewUserService(userRepo repository.UserRepoInterface, log *logrus.Logger, smsService sms_sender.SmsService) *UserService {
	return &UserService{userRepo: userRepo, log: log, smsService: smsService}
}

// Register User
// @Summary registering user
// @Description creating users
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.UserRegisterRequest true "Phone Number"
// @Router /users/register [post]
func (u *UserService) Register(c *gin.Context) {
	var req dto.UserRegisterRequest
	if err := c.BindJSON(&req); err != nil {
		u.log.Error("Invalid request")
		utils.ErrorResponse(c, 400, "Invalid request")
		return

	}
	phoneNumber, err := u.userRepo.GetPhoneNumberFromToken(req.VerifiedToken)
	if err != nil {
		u.log.Error("can not retrieve phone number from token")
		utils.ErrorResponse(c, 400, "can not retrieve phone number from token")
		return

	}
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		u.log.Error("can not hash password")
		utils.ErrorResponse(c, 500, "can not hash password")
		return
	}
	user := domain.User{
		Name:     req.Name,
		Password: passwordHash,
		Phone:    phoneNumber,
	}
	if err := u.userRepo.Register(&user); err != nil {
		u.log.Error(err.Error())
		utils.ErrorResponse(c, 500, err.Error())
		return
	}
	go u.smsService.SendSMS([]string{phoneNumber}, "welcome to our website")
	u.log.Info("user register successfully")
	utils.SuccessResponse(c, "user register successfully", dto.UserRegisterResponse{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.Phone,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
	})

}



// Login User
// @Summary login user
// @Description login users
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.UserLoginRequest true "Phone Number"
// @Router /users/login [post]
func (u *UserService) Login(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.BindJSON(&req); err != nil {
		u.log.Error("Invalid request")
		utils.ErrorResponse(c, 400, "Invalid request")
		return
	}
	user, err := u.userRepo.GetByPhoneNumber(req.PhoneNumber)
	if err != nil {
		u.log.Error("user not found")
		utils.ErrorResponse(c, 404, "user not found")
		return
	}
	if !utils.VerifyPassword(req.Password, user.Password) {
		u.log.Error("invalid password")
		utils.ErrorResponse(c, 400, "invalid password")
		return
	}
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Name, user.Role)
	if err != nil {
		u.log.Error("can not generate access token")
		utils.ErrorResponse(c, 500, "can not generate access token")
		return
	}
	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Name, user.Role)
	if err != nil {
		u.log.Error("can not generate refresh token")
		utils.ErrorResponse(c, 500, "can not generate refresh token")
		return
	}
	utils.SuccessResponse(c, "login successfully", dto.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}