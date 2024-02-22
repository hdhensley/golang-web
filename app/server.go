package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"golang-web/app/bootstrap"
	"golang-web/app/database"
	"golang-web/app/utils"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.DEBUG)
	utils.SetLogger(e.Logger)

	SetupRoutes(e)

	database.StartDB()
	database.AutoMigrate()

	utils.Logger.Fatal(e.Start(fmt.Sprintf(":%s", bootstrap.GetPort())))
}
