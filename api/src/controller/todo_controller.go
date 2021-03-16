package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/model"
)

type Controller struct{}

func (pc Controller) Index(c *gin.Context) {
	var s todo.Model
	todos, err := s.GetAll()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"todos": todos})
	}
}
