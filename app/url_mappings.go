package app

import (
	"persons/app/middleware"
	"persons/controllers"
)

func mapUrls() {
	router.GET("/persons", controller.FindEveryoneHandler)
	router.GET("/persons/:username", controller.FindPersonHandler)
	router.POST("/persons/login", controller.LoginHandler)
	router.POST("/persons", controller.CreatePersonHandler)
	router.PUT("/persons/:username", controller.ChangePersonHandler)
	router.DELETE("/persons/:username", controller.DeletePersonHandler)
}
