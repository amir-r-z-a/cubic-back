package server

import (
	"github.com/amir-r-z-a/cubic-back/config"
	"github.com/gin-gonic/gin"
)


func NewWebServer() *gin.Engine {

	server := gin.Default()
	
	return server
	
}

func Run(app *gin.Engine, appConfig *config.AppConfig) {
	app.Run(":" + appConfig.Port)
}