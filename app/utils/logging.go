package utils

import "github.com/labstack/echo/v4"

var Logger echo.Logger

func SetLogger(l echo.Logger) {
	Logger = l
}
