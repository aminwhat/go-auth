package main

import (
	"go-auth/db"
	"go-auth/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	routers.RegisterAuthRoutes(app, db)

	app.Run(":3000")
}
