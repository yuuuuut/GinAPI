package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/models"
)

type CommentController struct{}

var commentModel = new(models.Comment)

func (cr CommentController) Create(c *gin.Context) {
	comment, err := commentModel.CreateM(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(201, gin.H{"comment": comment})
	}
}
