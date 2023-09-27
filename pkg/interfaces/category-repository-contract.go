package interfaces

import "ecommerce/pkg/models"

type CategoryRepositoryContract interface {
	AddCategory(category models.Category) error
	GetAllCategory() (*[]models.Category, error)
}
