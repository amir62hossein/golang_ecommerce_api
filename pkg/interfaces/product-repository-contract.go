package interfaces

import "ecommerce/pkg/models"

type ProductRepositoryContract interface {
	AddProduct(models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProdcutByID(id string) (*models.Product, error)
	DeleteProduct(id string) error
	UpdateProduct(id string , product models.Product) error
}
