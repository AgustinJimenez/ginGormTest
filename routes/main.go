package routes

import (
	"go_practice/controllers"
	"go_practice/middlewares"

	"github.com/gin-gonic/gin"
)


func SetRoutes(r *gin.Engine) {
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/users", controllers.UsersIndex).Use(middlewares.Auth())
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
}