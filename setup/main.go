package setup

import (
	"go_practice/initializers"
	"go_practice/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)
func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.Request.Header.Get("Content-Type")
		if contentType != "application/json" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Request content type must be application/json",
			})
			return
		}

		c.Next()
	}
}

func GetApp() *gin.Engine {
	initializers.LoadEnvVariables()
	gin.SetMode(os.Getenv("GIN_MODE"))
	initializers.ConnectToDb()
	app := gin.Default()
	app.Use(JSONMiddleware())
	routes.SetRoutes(app)
  	return app
}