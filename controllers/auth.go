package controllers

import (
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

func (uc *AuthController) Signup(c *gin.Context) {
	var req dtos.AuthSignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"succeed": false, "meesage": "request body is not valid"})
		return
	}

	result, err := uc.AuthService.Signup(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

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
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
