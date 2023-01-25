package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"persons/app/middleware"
	"persons/database"
)

var router *gin.Engine

func StartApplication() {
	fmt.Println("Starting persons API")
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	db.Connection()
	RegisterMiddlewares()
	mapUrls()
	router.Run()
}

func RegisterMiddlewares() {
	router.Use(middleware.AuthenticateUser)
}
