package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println("Welcome To My Project mother Fucker")

	app := fiber.New()

	app.Listen("localhost:9000")

}
