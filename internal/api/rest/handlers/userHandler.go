package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-myobokucomerce-app/internal/api/rest"
	"go-myobokucomerce-app/internal/dto"
	"go-myobokucomerce-app/internal/repository"
	"go-myobokucomerce-app/internal/service"
	"net/http"
)

type userHandler struct {
	// waiting for svc userService
	svc service.UserService
}

func SetUpUserRoutes(rh *rest.RestHandler) {

	app := rh.App

	// Create instance of user service& inject to handler
	svc := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
	}
	handler := userHandler{
		svc: svc,
	}

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

	user := dto.UserSignup{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadGateway).JSON(fiber.Map{
			"Message": "di mohon masukan input data user yang benar.",
		})
	}

	token, err := h.svc.Signup(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "error dalam mendaftar data diri",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": token,
	})

}

func (h *userHandler) Login(ctx *fiber.Ctx) error {

	loginInput := dto.UserLogin{}
	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return ctx.Status(http.StatusBadGateway).JSON(fiber.Map{
			"Message": "di mohon masukan input data user yang benar.",
		})
	}

	token, err := h.svc.Login(loginInput.Email, loginInput.Password)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"Message": "error dalam login di mohon coba dengan data email dan password yang benar",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "your Success login",
		"token":   token,
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
