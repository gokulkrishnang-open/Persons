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
		r1.PUT("/:username", controller.ChangePersonHandler)
		r1.DELETE("/:username", controller.DeletePersonHandler)
	}
	r1.Use(middleware.AuthorizeUser)
	// router.GET("/persons", middleware.AuthorizeUser, controller.FindEveryoneHandler)
	// router.GET("/persons/:username", middleware.AuthorizeUser, controller.FindPersonHandler)
	router.POST("/persons/login", controller.LoginHandler)
	// router.POST("/persons", controller.CreatePersonHandler)
	// router.PUT("/persons/:username", middleware.AuthorizeUser, controller.ChangePersonHandler)
	// router.DELETE("/persons/:username", middleware.AuthorizeUser, controller.DeletePersonHandler)
}
