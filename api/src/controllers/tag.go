package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/models"
)

type TagController struct{}

var tagModel = new(models.Tag)

func (cr TagController) Index(c *gin.Context) {
	tags, err := tagModel.GetAll()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"tags": tags})
	}
}

func (cr TagController) Show(c *gin.Context) {
	id := c.Param("id")

	tag, err := tagModel.GetById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"tag": tag})
	}
}

func (cr TagController) Create(c *gin.Context) {
	err := tagModel.CreateM(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(201, gin.H{"message": "Create OK"})
	}
}
