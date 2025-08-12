package main

import (
	"go-auth/db"
	"go-auth/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	client := db.ConnectMongo("mongodb://localhost:27017")
	db := client.Database("go-auth")

	app := gin.Default()

	routers.RegisterAuthRoutes(app, db)

	app.Run(":3000")
}
