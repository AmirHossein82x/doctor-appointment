package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/AmirHossein82x/doctor-appointment/internal/app/constants"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/dto"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/ports"
	"github.com/AmirHossein82x/doctor-appointment/internal/app/utils"
	"github.com/AmirHossein82x/doctor-appointment/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	userRepo   ports.UserRepository
	log        *logrus.Logger
	smsService ports.SmsService
}

func NewUserService(userRepo ports.UserRepository, log *logrus.Logger, smsService ports.SmsService) *UserService {
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
	accessToken, err := utils.GenerateToken(user.ID, user.Name, user.Role, constants.TokenTypeAccess)
	if err != nil {
		u.log.Error("can not generate access token")
		utils.ErrorResponse(c, 500, "can not generate access token")
		return
	}
	refreshToken, err := utils.GenerateToken(user.ID, user.Name, user.Role, constants.TokenTypeRefresh)
	if err != nil {
		u.log.Error("can not generate refresh token")
		utils.ErrorResponse(c, 500, "can not generate refresh token")
		return
	}
	u.log.Info("user register successfully")
	utils.SuccessResponse(c, "user register successfully", dto.UserRegisterResponse{
		ID:           user.ID,
		Name:         user.Name,
		PhoneNumber:  user.Phone,
		Role:         user.Role,
		CreatedAt:    user.CreatedAt.Format(time.RFC3339),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
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
	accessToken, err := utils.GenerateToken(user.ID, user.Name, user.Role, constants.TokenTypeAccess)
	if err != nil {
		u.log.Error("can not generate access token")
		utils.ErrorResponse(c, 500, "can not generate access token")
		return
	}
	refreshToken, err := utils.GenerateToken(user.ID, user.Name, user.Role, constants.TokenTypeRefresh)
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

// verify access token
// @Summary verify access token
// @Description verify access token
// @Tags users
// @Accept json
// @Produce json
// @Router /users/verify-access-token [post]
func (u *UserService) VerifyAccessToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		u.log.Error("Empty access token")
		utils.ErrorResponse(c, 400, "Empty access token")
		return
	}
	// Extract the token from the header (format: "Bearer <token>")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.VerifyToken(tokenString, constants.TokenTypeAccess)
	if err != nil {
		u.log.Error("Invalid access token")
		utils.ErrorResponse(c, 400, "Invalid access token or token has expired")
		return
	}
	utils.SuccessResponse(c, "Valid access token", claims)
}

// get access token by refresh token
// @Summary get access token by refresh token
// @Description get access token by refresh token
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh Token"
// @Router /users/get-access-token-by-refresh-token [post]
func (u *UserService) GetAccessTokenByRefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.BindJSON(&req); err != nil {
		u.log.Error("Invalid request")
		utils.ErrorResponse(c, 400, "Invalid request")
		return
	}
	claims, err := utils.VerifyToken(req.RefreshToken, constants.TokenTypeRefresh)
	if err != nil {
		u.log.Error("Invalid refresh token")
		utils.ErrorResponse(c, 400, "Invalid refresh token or token has expired")
		return
	}
	accessToken, err := utils.GenerateToken(claims.ID, claims.Name, claims.Role, constants.TokenTypeAccess)
	if err != nil {
		u.log.Error("can not generate access token")
		utils.ErrorResponse(c, 500, "can not generate access token")
		return
	}
	utils.SuccessResponse(c, "access token generated successfully", dto.AccessTokenResponse{
		AccessToken: accessToken,
	})
}

// get forget password
// @Summary get forget password
// @Description get forget password
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.ForgetPasswordRequest true "Phone number"
// @Router /users/forget-password [post]
func (u *UserService) ForgetPassword(c *gin.Context) {
	var req dto.ForgetPasswordRequest
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

	// Encrypt UUID
	encryptedUUID, err := utils.EncryptUUID(user.ID)
	if err != nil {
		u.log.Error(err.Error())
		utils.ErrorResponse(c, 500, "internal server error")
		return
	}
	err = u.userRepo.SaveEncryptionKeyToRedis(encryptedUUID)
	if err != nil {
		u.log.Error(err.Error())
		utils.ErrorResponse(c, 500, "internal server error")
		return
	}
	// Generate password reset link
	resetLink := fmt.Sprintf("http://127.0.0.1:8080/forget-password?key=%s", encryptedUUID)
	go u.smsService.SendSMS([]string{req.PhoneNumber}, resetLink)
	utils.SuccessResponse(c, "reset password has send to you", nil)

}

// get reset password
// @Summary get reset password
// @Description get reset password
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.PasswordResetRequest true "password and password_retype"
// @Param key query string true "The encrypted key for password reset"
// @Router /users/reset-password [post]
func (u *UserService) ResetPassword(c *gin.Context) {
	encryptionKey := c.Query("key")
	var req dto.PasswordResetRequest
	if encryptionKey == "" {
		u.log.Error("bad request")
		utils.ErrorResponse(c, 400, "bad request")
		return

	}
	exists := u.userRepo.ExistsEncryptionKey(encryptionKey)
	if !exists {
		u.log.Error("your key is invalid or expired")
		utils.ErrorResponse(c, 400, "your key is invalid or expired")
		return
	}
	if err := c.BindJSON(&req); err != nil {
		u.log.Error("Invalid request")
		utils.ErrorResponse(c, 400, "Invalid request")
		return

	}
	if req.Password != req.PasswordRetype {
		u.log.Error("password and password retype are not matched")
		utils.ErrorResponse(c, 500, "password and password retype are not matched")
		return
	}

	userIdString, err := utils.DecryptUUID(encryptionKey)
	if err != nil {
		u.log.Error("can not decrypt key")
		utils.ErrorResponse(c, 500, "can not decrypt key")
		return

	}
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		u.log.Error("can not prase user id to uuid")
		utils.ErrorResponse(c, 500, "can not prase user id to uuid")
		return

	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		u.log.Error("can not hash password")
		utils.ErrorResponse(c, 500, "can not hash password")
		return

	}
	err = u.userRepo.UpdatePassword(userId, hashedPassword)

	if err != nil {
		u.log.Error("can not change the password")
		utils.ErrorResponse(c, 500, err.Error())
		return
	}
	err = u.userRepo.DeleteEncryptionKey(encryptionKey)
	if err != nil {
		u.log.Error("internal server error")
		utils.ErrorResponse(c, 500, "internal server error")
		return
	}
	utils.SuccessResponse(c, "password changes success fully", nil)

}
