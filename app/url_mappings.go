package app

import (
	"persons/controllers"
)

func mapUrls() {
	router.GET("/persons", controller.FindEveryoneHandler)
	router.GET("/persons/:username", controller.FindPersonHandler)
	router.POST("/persons", controller.CreatePersonHandler)
	router.PUT("/persons/:username", controller.ChangePersonHandler)
	router.DELETE("/persons/:username", controller.DeletePersonHandler)
}
