package handlers

import (
	"ecommerce/pkg/interfaces"
	"ecommerce/pkg/models"
	"ecommerce/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc interfaces.UserServiceContract
}

func NewUserHandler() *UserHandler {
	svc := services.NewUserService()
	return &UserHandler{
		svc: svc,
	}
}

func SetupUserRoutes(app *fiber.App) {
	handler := NewUserHandler()
	app.Post("/user/register", handler.Register)
	app.Post("/user/login", handler.Login)
}
func (h *UserHandler) Register(c *fiber.Ctx) error {

	user := new(models.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.svc.Register(*user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON("User Add Successfuly")
}
func (h *UserHandler) Login(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	foundUser, err := h.svc.Login(user.Email , user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(foundUser)
}
