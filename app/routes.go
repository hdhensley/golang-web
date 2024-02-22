package main

import (
	"github.com/labstack/echo/v4"
	"golang-web/app/controllers"
)

func SetupRoutes(e *echo.Echo) {
	e.Static("/public", "public")

	e.GET("/", controllers.Home)
	e.GET("/hello/user", controllers.UserTest)
	e.GET("/second", controllers.SecondPage)

	e.GET("/hello", controllers.Hello)
}
