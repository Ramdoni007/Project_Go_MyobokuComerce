package api

import (
	"github.com/gofiber/fiber/v2"
	"go-myobokucomerce-app/config"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Listen(config.ServerPort)
}
