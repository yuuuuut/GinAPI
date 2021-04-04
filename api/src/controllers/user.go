package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/models"
)

type UserController struct{}

var (
	userModel = new(models.User)
)

func (cr UserController) Show(c *gin.Context) {
	id := c.Param("id")

	user, err := userModel.GetById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"user": user})
	}
}

func (cr UserController) Create(c *gin.Context) {
	user, userErr := userModel.CreateM(c)
	profile, profileErr := profileModel.CreateM(c, user.ID)

	if userErr != nil {
		c.JSON(400, gin.H{"error": userErr.Error()})
		return
	} else if profileErr != nil {
		c.JSON(400, gin.H{"error": profileErr.Error()})
		return
	} else {
		c.JSON(201, gin.H{"user": user, "profile": profile})
	}
}
