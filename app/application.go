package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"persons/database"
)

var router *gin.Engine

func StartApplication() {
	fmt.Println("Starting persons API")
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	db.Connection()
	mapUrls()
	router.Run()
}
