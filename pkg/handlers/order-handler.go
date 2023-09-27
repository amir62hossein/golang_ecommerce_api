package handlers

import (
	"ecommerce/pkg/interfaces"
	"ecommerce/pkg/models"
	"ecommerce/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	svc interfaces.OrderServiceContract
}

func NewOrderHandler() *OrderHandler {
	svc := services.NewOrderService()
	return &OrderHandler{
		svc: svc,
	}
}

func SetupOrderRoutes(app *fiber.App) {
	h := NewOrderHandler()
	app.Post("/order", h.AddNewOrder)
}
func (h *OrderHandler) AddNewOrder(c *fiber.Ctx) error {

	order := new(models.Orders)

	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.svc.AddOrder(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON("order added successfuly")
}
