package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/controllers"
	"github.com/yuuuuut/gin-api/src/middleware"
)

func Init() {
	r := Router()
	r.Run()
}

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Loging)

	todos := r.Group("/todos")
	{
		ctrl := new(controllers.TodoController)
		todos.GET("", middleware.FirebaseAuth, ctrl.Index)
		todos.GET("/:id", ctrl.Show)
		todos.POST("", ctrl.Create)
		todos.PUT("/:id", ctrl.Update)
		todos.DELETE("/:id", ctrl.Delete)
	}

	return r
}
