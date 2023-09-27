package interfaces

import "ecommerce/pkg/models"

type CategoryServiceContract interface {
	AddCategory(category models.Category) error
	GetAllCategory() (*[]models.Category, error)
}
