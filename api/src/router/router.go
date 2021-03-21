package router

import (
	"github.com/gin-gonic/gin"
	todo "github.com/yuuuuut/gin-api/src/controller"
	"github.com/yuuuuut/gin-api/src/middleware"
)

func Init() {
	r := Router()
	r.Run()
}

func Router() *gin.Engine {
	r := gin.Default()

	todos := r.Group("/todos")
	{
		ctrl := todo.Controller{}
		todos.GET("", middleware.FirebaseAuth, ctrl.Index)
		todos.GET("/:id", ctrl.Show)
		todos.POST("", ctrl.Create)
		todos.PUT("/:id", ctrl.Update)
		todos.DELETE("/:id", ctrl.Delete)
	}

	return r
}
