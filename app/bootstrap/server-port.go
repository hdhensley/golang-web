package bootstrap

import (
	"os"
)

func GetPort() string {
	port := os.Getenv("PORT")

	if "" == port {
		return "3005"
	}

	return port
}
