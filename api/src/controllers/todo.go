package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/models"
)

type TodoController struct{}

var todoModel = new(models.Todo)

func (cr TodoController) Index(c *gin.Context) {
	var (
		offset, limit         string
		offsetBool, limitBool bool
	)

	offset, offsetBool = c.GetQuery("offset")
	if !offsetBool {
		offset = "0"
	}
	limit, limitBool = c.GetQuery("limit")
	if !limitBool {
		limit = "3"
	}

	todos, err := todoModel.GetAll(offset, limit)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"todos": todos})
	}
}

func (cr TodoController) Show(c *gin.Context) {
	id := c.Param("id")

	todo, err := todoModel.GetById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"todo": todo})
	}
}

func (cr TodoController) Create(c *gin.Context) {
	todo, errorMessages, err := todoModel.CreateM(c)
	if errorMessages != nil {
		c.JSON(400, gin.H{"errorMessages": errorMessages})
		return
	} else if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(201, gin.H{"todo": todo})
	}
}

func (cr TodoController) Update(c *gin.Context) {
	id := c.Param("id")

	todo, err := todoModel.UpdateById(id, c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"todo": todo})
	}
}

func (cr TodoController) Delete(c *gin.Context) {
	id := c.Param("id")

	todo, err := todoModel.DeleteById(id)
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(204, gin.H{"todo": todo})
	}
}
