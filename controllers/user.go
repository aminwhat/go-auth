package controllers

import (
	"fmt"
	"go-auth/dtos"
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

// @Summary Get All Users
// @Description Get all users with pagination and optional phone number search
// @Security ApiKeyAuth
// @Accept json
// @Param page query int false "Page number (default: 1)" minimum(1)
// @Param pageSize query int false "Page size (default: 10)" minimum(1) maximum(100)
// @Param phone query string false "Phone number search (partial match)"
// @Produce json
// @Success 200 {object} dtos.GetAllUsersResponse
// @Failure 400 {object} controllers_swagger.GetAllUsersBadResponse "Bad Request"
// @Router /user/all [get]
func (uc *UserController) GetAllUsers(c *gin.Context) {
	var request dtos.GetAllUsersRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"succeed": false,
			"message": err.Error(),
		})
		return
	}

	result, err := uc.UserService.GetAllUsersWithPagination(request)

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

	if err != nil || !result.Succeed {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
