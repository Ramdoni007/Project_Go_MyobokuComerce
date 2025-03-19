package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-myobokucomerce-app/internal/api/rest"
	"net/http"
)

type userHandler struct {
	// svc userService

}

func SetUpUserRoutes(rh *rest.RestHandler) {

	app := rh.App

	// Create instance of user service& injectto handler
	handler := userHandler{}

	//Public EndPoint
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	//Private Endpoint
	app.Get("/verify", handler.GetVerification)
	app.Post("/verify", handler.Verify)
	app.Post("/profile", handler.CreateProfile)
	app.Get("/profile", handler.GetProfile)

	app.Post("/cart", handler.AddToCart)
	app.Get("/cart", handler.GetCart)
	app.Get("/order", handler.GetOrders)
	app.Get("/order/:id", handler.GetOrder)

	app.Post("/become-seller", handler.BecomeSeller)

}

func (h *userHandler) Register(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "your success register",
	})

}

func (h *userHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "your Success login",
	})
}

func (h *userHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "your Success Verify",
	})
}

func (h *userHandler) GetVerification(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Get Verification",
	})
}

func (h *userHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "CreateProfile",
	})
}

func (h *userHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Get Profile",
	})
}

func (h *userHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "add to cart",
	})
}

func (h *userHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "get cart",
	})
}

func (h *userHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "get orders",
	})
}

func (h *userHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "get order by id",
	})
}

func (h *userHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "become seller",
	})
}
