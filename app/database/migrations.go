package database

import (
	"golang-web/app/models"
	"golang-web/app/utils"
	"os"
)

func AutoMigrate() {
	// Automigrate unless the environment variable DB_AUTOMIGRATE is set to false
	utils.Logger.Info("Auto migration started")
	if os.Getenv("DB_AUTOMIGRATE") == "false" {
		utils.Logger.Info("Skipping auto migration")
		return
	}

	db := getDBPool()

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	utils.Logger.Info("Auto migration finished")
}
