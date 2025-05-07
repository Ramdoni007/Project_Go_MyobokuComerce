package handlers

import (
	"go-myobokucomerce-app/internal/api/rest"
	"go-myobokucomerce-app/internal/dto"
	"go-myobokucomerce-app/internal/repository"
	"go-myobokucomerce-app/internal/service"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	svc service.UserService
}

func SetUpUserRoutes(rh *rest.RestHandler) {

	app := rh.App

	// Create instance of user service& inject to handler
	svc := service.UserService{
		Repo:   repository.NewUserRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}
	handler := userHandler{
		svc: svc,
	}

	pubRoutes := app.Group("/users")

	//Public EndPoint
	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)

	//Private Endpoint
	pvtRoutes.Get("/verify", handler.GetVerification)
	pvtRoutes.Post("/verify", handler.Verify)
	pvtRoutes.Post("/profile", handler.CreateProfile)
	pvtRoutes.Get("/profile", handler.GetProfile)

	pvtRoutes.Post("/cart", handler.AddToCart)
	pvtRoutes.Get("/cart", handler.GetCart)
	pvtRoutes.Get("/order", handler.GetOrders)
	pvtRoutes.Get("/order/:id", handler.GetOrder)

	pvtRoutes.Post("/become-seller", handler.BecomeSeller)

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
		"message": "register",
		"token":   token,
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

	user := h.svc.Auth.GetCurrentUser(ctx)

	//request
	var req dto.VerificationInput

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "please provide a valid input",
		})
	}

	err := h.svc.VerifyCode(user.ID, req.Code)

	if err != nil {
		log.Printf("%v", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Verified successfully",
	})
}

func (h *userHandler) GetVerification(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)
	log.Println(user)

	//create verification code and update to user profile in DB
	err := h.svc.GetVerificationCode(user)
	log.Println(err)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "unable to generate verification code ",
			"data":    err,
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "get Verification",
	})
}

func (h *userHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "CreateProfile",
	})
}

func (h *userHandler) GetProfile(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)
	log.Println(user)

	// call user service and perform get profile
	profile, err := h.svc.GetProfile(user.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to get profile",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get profile",
		"profile": profile,
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
