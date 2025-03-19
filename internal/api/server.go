package api

import (
	"github.com/gofiber/fiber/v2"
	"go-myobokucomerce-app/config"
	"go-myobokucomerce-app/internal/api/rest"
	"go-myobokucomerce-app/internal/api/rest/handlers"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	rh := &rest.RestHandler{
		App: app,
	}

	setUpRoutes(rh)

	app.Listen(config.ServerPort)
}

func setUpRoutes(rh *rest.RestHandler) {

	// user handlers
	handlers.SetUpUserRoutes(rh)

	//transactions coming
	// catalog coming

}
