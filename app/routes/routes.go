package routes

import (
	"alimrndev/super-admin-golang-backend/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userController := controllers.UserController{}

	// User Routes
	router.GET("/", userController.Index)
	router.GET("/users", userController.GetAll)
	router.GET("/users/:id", userController.GetByID)
	router.POST("/users", userController.Create)
	router.PUT("/users/:id", userController.Update)
	router.DELETE("/users/:id", userController.Delete)
}
