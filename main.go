package main

import (
	"go-auth/db"
	"go-auth/routers"
	"log"
	"os"

	docs "go-auth/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := db.ConnectMongo(
		os.Getenv(("MONGO_URI")),
		os.Getenv("MONGO_USERNAME"),
		os.Getenv("MONGO_PASSWORD"),
	)
	db := client.Database(os.Getenv("MONGO_DB_NAME"))

	app := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	routers.RegisterAuthRoutes(app, db)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	app.Run(":3000")
}
