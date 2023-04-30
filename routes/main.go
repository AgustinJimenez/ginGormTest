package routes

import (
	"go_practice/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
}