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

	if err != nil || !result.Succeed {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Get User By Id
// @Description Get A User Data using Specified Id
// @Security ApiKeyAuth
// @Accept json
// @Param userId path string true "User ID"
// @Produce json
// @Success 200 {object} dtos.GetCurrentUserResponse
// @Failure 400 {object} controllers_swagger.GetCurrentUserBadResponse "Something Unknown Happend"
// @Router /user/{userId} [get]
func (uc *UserController) GetUserById(c *gin.Context) {
	userId := c.Param("userId")

	result, err := uc.UserService.GetUser(userId)

	if err != nil || result.Succeed {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
