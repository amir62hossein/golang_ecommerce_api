package services

import (
	"ecommerce/pkg/interfaces"
	"ecommerce/pkg/models"
	"ecommerce/pkg/repository"
)

type CategoryService struct {
	repository interfaces.CategoryRepositoryContract
}

func NewCategoryService() *CategoryService {
	repo := repository.NewPostgres()
	return &CategoryService{
		repository: repo,
	}
}
func (r *CategoryService) AddCategory(category models.Category) error {
	if err := r.repository.AddCategory(category); err != nil {
		return err
	}
	return nil
}
func (r *CategoryService) GetAllCategory() (*[]models.Category, error) {
	categories, err := r.repository.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
