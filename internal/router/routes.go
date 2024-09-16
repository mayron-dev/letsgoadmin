package router

import (
	"github.com/gin-gonic/gin"
)

func registerApiRoutes(apiRouter *gin.RouterGroup) {

	apiRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
