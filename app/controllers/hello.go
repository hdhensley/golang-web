package controllers

import (
	"github.com/labstack/echo/v4"
	"golang-web/app/utils"
	components "golang-web/app/views/components"
	"net/http"
)

func Hello(c echo.Context) error {
	return utils.Render(c, http.StatusOK, components.Hello())
}
