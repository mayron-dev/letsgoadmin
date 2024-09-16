package main

import (
	"github.com/mayron-dev/letsgoadmin/config"
	"github.com/mayron-dev/letsgoadmin/internal/router"
)

func main() {
	logger := config.GetLogger("main")
	err := config.InitConfig()
	if err != nil {
		logger.Errorf("config error: %s", err.Error())
		panic(err)
	}
	router.InitRouter()
}
