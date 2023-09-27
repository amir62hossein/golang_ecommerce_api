package interfaces

import "ecommerce/pkg/models"

type ProductServiceContract interface {
	AddProduct(models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProductById(id string) (*models.Product, error)
	DeleteProduct(id string) error
	UpdateProduct(id string, product models.Product) error
}
