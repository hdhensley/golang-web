package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang-web/app/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

var Db *gorm.DB

func getDBPool() *gorm.DB {
	return Db
}

func StartDB() {
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Warnf("No .env file found, falling back to os env, %s", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	utils.Logger.Info("Database connected")

	Db = db

	//utils.Logger.Info("Migration started")
	//AutoMigrate()
	//utils.Logger.Info("Migration finished")
}
