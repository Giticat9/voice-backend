package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"voice/main/internal/api"
	"voice/main/internal/common/config"
)

func main() {
	config.LoadConfig()

	ginMode := config.AppConfig.GinMode
	gin.SetMode(ginMode)

	host := config.AppConfig.ServerHost
	port := config.AppConfig.ServerPort
	addr := host + ":" + port

	router := api.CreateAppRouter()
	router.InitAppRoutes()

	err := router.Run(addr)

	if err != nil {
		log.Fatal(err)
	}
}
