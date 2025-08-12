package controllers

import (
	"fmt"
	"go-auth/services"
	"net/http"

	"go-auth/dtos"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

// @Summary Signup Method
// @Description Sends an OTP code to the phoneNumber (fake OTP, just printed in the console)
// @Accept json
// @Produce json
// @Param body body dtos.AuthSignupRequest true "AuthSignupRequest model"
// @Success 200 {object} dtos.AuthSignupResponse
// @Failure 400 {object} controllers_swagger.AuthSignupBadResponse "Request body is not valid"
// @Router /auth/signup [post]
func (uc *AuthController) Signup(c *gin.Context) {
	var req dtos.AuthSignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"succeed": false, "meesage": "request body is not valid"})
		return
	}

	result, err := uc.AuthService.Signup(req)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Signup Confirm Otp Method
// @Description Validate the PhoneNumber with the OtpCode alongside withit
// @Accept json
// @Produce json
// @Param body body dtos.AuthSignupConfirmOtpRequest true "AuthSignupConfirmOtpRequest model"
// @Success 200 {object} dtos.AuthTokenResponse
// @Failure 400 {object} controllers_swagger.AuthSignupConfirmOtpBadResponse "Request body is not valid"
// @Router /auth/signup/confirm-otp [post]
func (uc *AuthController) SignupConfirmOtp(c *gin.Context) {
	var req dtos.AuthSignupConfirmOtpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"succeed": false, "meesage": "request body is not valid"})
		return
	}

	result, err := uc.AuthService.SignupConfirmOtp(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (uc *AuthController) Login(c *gin.Context) {
	var req dtos.AuthLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"succeed": false, "meesage": "request body is not valid"})
		return
	}

	result, err := uc.AuthService.Login(req)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
