package main

import (
	"alimrndev/superadmin-golang-backend/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router)

	router.Run(":8080")
}
