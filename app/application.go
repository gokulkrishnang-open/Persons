package app

import (
	"fmt"

	"persons/config"
	db "persons/database"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func StartApplication() {
	fmt.Println("Starting persons API")
	gin.SetMode(gin.ReleaseMode)
	Init()
	router = gin.Default()
	db.Connection()
	mapUrls()
	router.Run(":" + config.AppConfig.Port)
}

func Init() {
	config.LoadConfigs()
}
