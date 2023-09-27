package handlers

import (
	"ecommerce/pkg/interfaces"
	"ecommerce/pkg/models"
	"ecommerce/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type CommentHandler struct {
	svc interfaces.CommentServiceContract
}

func NewCommentHandler() *CommentHandler {
	svc := services.NewCommentService()
	return &CommentHandler{
		svc: svc,
	}
}

func SetupCommentRoutes(app *fiber.App) {
	h := NewCommentHandler()
	app.Post("/comment", h.AddComment)
	app.Get("/comment", h.GetAllComment)
	app.Get("/comment/:id", h.GetCommentById)
	app.Delete("/comment/:id", h.DeleteComment)
	app.Get("/comment/product-id/:id", h.GetCommentByProductId)
	app.Get("/comment/user-id/:id", h.GetCommentByUserId)
}

func (h *CommentHandler) AddComment(c *fiber.Ctx) error {
	comment := new(models.Comments)

	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.svc.AddComment(*comment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON("comment added")
}
func (h *CommentHandler) GetAllComment(c *fiber.Ctx) error {
	comments, err := h.svc.GetAllComments()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(comments)
}
func (h *CommentHandler) GetCommentById(c *fiber.Ctx) error {
	id := c.Params("id")
	comment, err := h.svc.GetCommentByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(comment)
}
func (h *CommentHandler) DeleteComment(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.svc.DeleteComment(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("comment deleted")
}
func (h *CommentHandler) GetCommentByProductId(c *fiber.Ctx) error {

	id := c.Params("id")

	comments, err := h.svc.FindCommentByProductID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(comments)
}
func (h *CommentHandler) GetCommentByUserId(c *fiber.Ctx) error {
	id := c.Params("id")

	comments, err := h.svc.FindCommentByUserID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(comments)
}
