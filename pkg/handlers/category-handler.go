package handlers

import (
	"ecommerce/pkg/interfaces"
	"ecommerce/pkg/models"
	"ecommerce/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	svc interfaces.CategoryServiceContract
}

func NewCategoryHandler() *CategoryHandler {
	svc := services.NewCategoryService()
	return &CategoryHandler{
		svc: svc,
	}
}

func SetupCategoryRoutes(app *fiber.App) {
	h := NewCategoryHandler()
	app.Post("/category", h.AddCategory)
	app.Get("/category", h.GetAllCategories)
}
func (h *CategoryHandler) AddCategory(c *fiber.Ctx) error {

	var category models.Category

	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.svc.AddCategory(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("category added")
}
func (h *CategoryHandler) GetAllCategories(c *fiber.Ctx) error {
	categories, err := h.svc.GetAllCategory()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(categories)
}
