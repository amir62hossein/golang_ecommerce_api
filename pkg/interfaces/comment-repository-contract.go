package interfaces

import "ecommerce/pkg/models"

type CommentRepositoryContract interface {
	AddComment(comment *models.Comments) error
	GetAllComments() (*[]models.Comments, error)
	GetCommentByID(id string) (*models.Comments, error)
	DeleteComment(id string) error
	FindCommentByProductID(id string) (*[]models.Comments, error)
	FindCommentByUserID(id string) (*[]models.Comments, error)
}
