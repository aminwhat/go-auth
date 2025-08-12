package controllers

import (
	"fmt"
	"go-auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// @Summary Current User
// @Description Get Current User Data using Jwt Token
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} dtos.GetCurrentUserResponse
// @Failure 400 {object} controllers_swagger.GetCurrentUserBadResponse "Something Unknown Happend"
// @Router /user [get]
func (uc *UserController) GetCurrentUser(c *gin.Context) {
	userId, _ := c.Get("userId")
	fmt.Println("UserId :" + userId.(string))

	result, err := uc.UserService.GetUser(userId.(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
