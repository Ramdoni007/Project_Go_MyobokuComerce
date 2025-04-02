package api

import (
	"github.com/gofiber/fiber/v2"
	"go-myobokucomerce-app/config"
	"go-myobokucomerce-app/internal/api/rest"
	"go-myobokucomerce-app/internal/api/rest/handlers"
	"go-myobokucomerce-app/internal/domain"
	"go-myobokucomerce-app/internal/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error %v\n", err)
	}
	log.Println("database connected")

	//run Migration
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		return
	}

	auth := helper.SetUpAuth(config.AppSecret)

	rh := &rest.RestHandler{
		App:  app,
		DB:   db,
		Auth: auth,
	}

	setUpRoutes(rh)

	err = app.Listen(config.ServerPort)
	if err != nil {
		return
	}
}

func setUpRoutes(rh *rest.RestHandler) {

	// user handlers
	handlers.SetUpUserRoutes(rh)

	//transactions coming
	// catalog coming

}
