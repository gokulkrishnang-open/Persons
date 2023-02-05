package app

import (
	"persons/app/middleware"
	"persons/controllers"
)

func mapUrls() {
	router.GET("/persons", middleware.AuthorizeUser, controller.FindEveryoneHandler)
	router.GET("/persons/:username", middleware.AuthorizeUser, controller.FindPersonHandler)
	router.POST("/persons/login", controller.LoginHandler)
	router.POST("/persons", controller.CreatePersonHandler)
	router.PUT("/persons/:username", middleware.AuthorizeUser, controller.ChangePersonHandler)
	router.DELETE("/persons/:username", middleware.AuthorizeUser, controller.DeletePersonHandler)
}
