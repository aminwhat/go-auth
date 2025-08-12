package routers

import (
	"go-auth/controllers"
	"go-auth/middlewares"
	"go-auth/repositories"
	"go-auth/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterUserRoutes(r *gin.Engine, db *mongo.Database) {
	jwtService := services.NewJwtService("my_very_secret_key")

	userRepo := repositories.NewUserRepository(db)

	userService := services.NewUserService(userRepo)

	userController := controllers.NewUserController(userService)

	user := r.Group("/user")

	user.Use(middlewares.AuthMiddleware(jwtService))

	user.GET("/", userController.GetCurrentUser)
	user.GET("/:userId", userController.GetUserById)
}
