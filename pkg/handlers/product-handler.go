package handlers

import (
	"ecommerce/pkg/interfaces"
	"ecommerce/pkg/middleware"
	"ecommerce/pkg/models"
	"ecommerce/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	svc interfaces.ProductServiceContract
}

func NewProductHandler() *ProductHandler {
	svc := services.NewProductService()
	return &ProductHandler{
		svc: svc,
	}
}

func SetupProductRoutes(app *fiber.App) {
	h := NewProductHandler()
	app.Get("/products", middleware.Auth(), h.GetAllProducts)
	app.Get("/products/:id", h.GetProdcutByID)
	app.Post("/products",middleware.Admin(), h.AddProduct)
	app.Delete("/products/:id", h.DeleteProduct)
	app.Put("/products/:id", h.UpdateProduct)
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {

	products, err := h.svc.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(products)
}
func (h *ProductHandler) AddProduct(c *fiber.Ctx) error {

	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.svc.AddProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON("product added")
}
func (h *ProductHandler) GetProdcutByID(c *fiber.Ctx) error {

	id := c.Params("id")

	product, err := h.svc.GetProductById(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(product)
}
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.svc.DeleteProduct(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON("product deleted")
}
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.svc.UpdateProduct(id, product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("product update")
}
