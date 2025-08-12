package routers

import (
	"go-auth/controllers"
	"go-auth/repositories"
	"go-auth/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterAuthRoutes(r *gin.Engine, db *mongo.Database) {

	// Initialize user Repository
	userRepo := repositories.NewUserRepository(db)

	// Initialize service with database instance
	authService := services.NewAuthService(userRepo)

	// Initialize controller with service
	userController := controllers.NewAuthController(authService)

	user := r.Group("/user")

	user.GET("/signup", userController.Signup)
	user.POST("/login", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"msg": "users"}) })

}
