package controllers

import (
	"github.com/labstack/echo/v4"
	"golang-web/app/utils"
	views "golang-web/app/views/pages"
	"net/http"
)

func Home(c echo.Context) error {
	return utils.Render(c, http.StatusOK, views.Home())
}
