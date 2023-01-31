package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "persons/app/middleware"
	"persons/config"
	"persons/database"
)

var router *gin.Engine

func StartApplication() {
	fmt.Println("Starting persons API")
	gin.SetMode(gin.ReleaseMode)
	Init()
	router = gin.Default()
	db.Connection()
	// RegisterMiddlewares()
	mapUrls()
	router.Run(":" + config.AppConfig.Port)
}

func Init() {
	config.LoadConfigs()
}

// func RegisterMiddlewares() {
// 	router.Use(middleware.AuthenticateUser)
// }
