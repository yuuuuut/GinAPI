package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/models"
)

type TagController struct{}

var tagModel = new(models.Tag)

func (cr TagController) Create(c *gin.Context) {
	err := tagModel.CreateM(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(201, gin.H{"message": "Create OK"})
	}
}
