package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Config struct {
	Port string
}

func LoadEnv(l echo.Logger) {
	err := godotenv.Load()
	if err != nil {
		l.Info("No .env file found")
		return
	}
}
