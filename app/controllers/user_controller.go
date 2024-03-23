package controllers

import (
	"alimrndev/superadmin-golang-backend/app/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func (uc *UserController) GetAll(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(200, users)
}

func (uc *UserController) GetByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	models.DB.First(&user, id)
	c.JSON(200, user)
}

func (uc *UserController) Create(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	models.DB.Create(&user)
	c.JSON(200, user)
}

func (uc *UserController) Update(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	models.DB.First(&user, id)
	c.BindJSON(&user)
	models.DB.Save(&user)
	c.JSON(200, user)
}

func (uc *UserController) Delete(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	models.DB.First(&user, id)
	models.DB.Delete(&user)
	c.JSON(200, gin.H{"message": "success"})
}
