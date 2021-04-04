package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/models"
)

type ProfileController struct{}

var (
	profileModel = new(models.Profile)
)

func (cr ProfileController) Show(c *gin.Context) {
	id := c.Param("id")

	profile, err := profileModel.GetById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"profile": profile})
	}
}
