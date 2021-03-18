package router

import (
	"github.com/gin-gonic/gin"
	todo "github.com/yuuuuut/gin-api/src/controller"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	todos := r.Group("/todos")
	{
		ctrl := todo.Controller{}
		todos.GET("", ctrl.Index)
		todos.GET("/:id", ctrl.Show)
		todos.POST("", ctrl.Create)
	}

	return r
}
