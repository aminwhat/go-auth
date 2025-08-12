package controllers

import (
	"go-auth/services"
	"net/http"

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
	result := uc.AuthService.Signup()
	c.JSON(http.StatusOK, result)
}

func (uc *AuthController) Login(c *gin.Context) {
	result := uc.AuthService.Login()
	c.JSON(http.StatusOK, result)
}
