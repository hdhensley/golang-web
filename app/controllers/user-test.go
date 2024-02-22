package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"golang-web/app/models"
	"golang-web/app/utils"
	components "golang-web/app/views/components"
	"gorm.io/gorm"
	"net/http"
)

func UserTest(c echo.Context) error {
	user := models.User{
		Model:        gorm.Model{},
		Name:         "testUsername",
		Email:        nil,
		Age:          0,
		Birthday:     nil,
		MemberNumber: sql.NullString{},
	}
	return utils.Render(c, http.StatusOK, components.UserTest(user))
}
