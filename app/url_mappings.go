package app

import (
	"persons/app/middleware"
	controller "persons/controllers"
)

func mapUrls() {
	r1 := router.Group("/persons")
	{
		r1.GET("/", controller.FindEveryoneHandler)
		r1.GET("/:username", controller.FindPersonHandler)
		r1.POST("/", controller.CreatePersonHandler)
		r1.PUT("/:username", controller.UpdatePersonHandler)
		r1.DELETE("/:username", controller.DeletePersonHandler)
	}
	r1.Use(middleware.AuthorizeUser)
	router.POST("/persons/login", controller.LoginHandler)
}
