package routers

import (
	"go-auth/controllers"
	"go-auth/repositories"
	"go-auth/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterAuthRoutes(r *gin.Engine, db *mongo.Database) {

	jwtService := services.NewJwtService("my_very_secret_key")

	// Initialize user Repository
	userRepo := repositories.NewUserRepository(db)

	// Initialize auth register repository
	authRegisterRepo := repositories.NewAuthRegisterRepository(db)

	// Initialize service with database instance
	authService := services.NewAuthService(userRepo, authRegisterRepo, jwtService)

	// Initialize controller with service
	authController := controllers.NewAuthController(authService)

	auth := r.Group("/auth")

	auth.POST("/signup", authController.Signup)
	auth.POST("/signup/confirm-otp", authController.SignupConfirmOtp)

}
