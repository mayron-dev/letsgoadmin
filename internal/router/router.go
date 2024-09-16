package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mayron-dev/letsgoadmin/config"
)

var (
	logger *config.Logger
)

func InitRouter() error {
	logger = config.GetLogger("Router")
	env := config.GetEnv()
	if env.ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	apiGroup := r.Group("/api")
	registerApiRoutes(apiGroup)

	r.Static("/assets", "./web/dist/assets")
	// Servir index.html para as demais rotas
	r.NoRoute(func(c *gin.Context) {
		c.File("./web/dist/index.html")
	})
	logger.Info("starting server...")
	r.Run(fmt.Sprintf(":%s", "8080"))
	return nil
}
