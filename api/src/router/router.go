package router

import (
	"github.com/gin-gonic/gin"
	todo "github.com/yuuuuut/gin-api/src/controller"
	"github.com/yuuuuut/gin-api/src/middleware"
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
		todos.GET("", middleware.FirebaseAuth, ctrl.Index)
		todos.GET("/:id", ctrl.Show)
		todos.POST("", ctrl.Create)
		todos.PUT("/:id", ctrl.Update)
	}

	return r
}
