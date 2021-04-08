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
		todos.POST("", middleware.FirebaseAuth, ctrl.Create)
		todos.PUT("/:id", ctrl.Update)
		todos.DELETE("/:id", ctrl.Delete)
	}

	users := r.Group("/users")
	{
		ctrl := new(controllers.UserController)
		users.GET("/:id", ctrl.Show)
		users.POST("", ctrl.Create)
	}

	profiles := r.Group("/profiles")
	{
		ctrl := new(controllers.ProfileController)
		profiles.GET("/:id", middleware.FirebaseAuth, ctrl.Show)
	}

	tags := r.Group("/tags")
	{
		ctrl := new(controllers.TagController)
		tags.GET("/:id", ctrl.Show)
		tags.POST("", ctrl.Create)
	}

	return r
}
