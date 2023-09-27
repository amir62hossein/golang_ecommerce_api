package services

import (
	"ecommerce/pkg/interfaces"
	"ecommerce/pkg/models"
	"ecommerce/pkg/repository"
)

type CommentService struct {
	repo interfaces.CommentRepositoryContract
}

func NewCommentService() *CommentService {
	repo := repository.NewPostgres()
	return &CommentService{
		repo: repo,
	}
}
func (r *CommentService) AddComment(comment models.Comments) error {
	if err := r.repo.AddComment(&comment); err != nil {
		return err
	}
	return nil
}
func (r *CommentService) GetAllComments() (*[]models.Comments, error) {
	comments, err := r.repo.GetAllComments()
	if err != nil {
		return nil, err
	}
	return comments, nil
}
func (r *CommentService) GetCommentByID(id string) (*models.Comments, error) {
	comment, err := r.repo.GetCommentByID(id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
func (r *CommentService) DeleteComment(id string) error {
	if err := r.repo.DeleteComment(id); err != nil {
		return err
	}
	return nil
}
func (r *CommentService) FindCommentByProductID(id string) (*[]models.Comments, error) {
	comments, err := r.repo.FindCommentByProductID(id)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
func (r *CommentService) FindCommentByUserID(id string) (*[]models.Comments, error) {
	comments, err := r.repo.FindCommentByUserID(id)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
