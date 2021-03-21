package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/helper"
	"github.com/yuuuuut/gin-api/src/model"
)

type Controller struct{}

func (pc Controller) Index(c *gin.Context) {
	var m todo.Model

	offset := helper.GetQueryIndex(c, "offset")
	limit := helper.GetQueryIndex(c, "limit")

	todos, err := m.GetAll(offset, limit)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"todos": todos})
	}
}

func (pc Controller) Show(c *gin.Context) {
	var m todo.Model
	id := c.Param("id")

	todo, err := m.GetById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"todo": todo})
	}
}

func (pc Controller) Create(c *gin.Context) {
	var m todo.Model

	todo, errorMessages, err := m.CreateM(c)
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

func (pc Controller) Update(c *gin.Context) {
	var m todo.Model
	id := c.Param("id")

	todo, err := m.UpdateById(id, c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"todo": todo})
	}
}

func (pc Controller) Delete(c *gin.Context) {
	var m todo.Model
	id := c.Param("id")

	todo, err := m.DeleteById(id)
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(204, gin.H{"todo": todo})
	}
}
