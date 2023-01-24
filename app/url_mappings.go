package app

import (
	"fmt"
	"persons/controllers"
)

func mapUrls() {
	fmt.Println("mapping urls")
	router.GET("/persons", controller.FindEveryoneHandler)
	router.GET("/persons/:username", controller.FindPersonHandler)
	router.POST("/persons", controller.CreatePersonHandler)
	router.PUT("/persons/:username", controller.ChangePersonHandler)
	router.DELETE("/persons/:username", controller.DeletePersonHandler)
}
