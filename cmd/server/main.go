package main

import (
	"log"

	"ginLearn/internal/app"
	"ginLearn/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	application := app.NewApp()

	r := gin.Default()
	// register routes
	routes.RegisterUserRoutes(r, application)

	log.Println("Server running on :8000")
	r.Run(":8080")
}
