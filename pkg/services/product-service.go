package services

import (
	"ecommerce/pkg/interfaces"
	"ecommerce/pkg/models"
	"ecommerce/pkg/repository"
)

type ProductService struct {
	repository interfaces.ProductRepositoryContract
}

func NewProductService() *ProductService {
	repo := repository.NewPostgres()
	return &ProductService{
		repository: repo,
	}
}

func (svc *ProductService) GetAllProducts() ([]models.Product, error) {
	products, err := svc.repository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (svc *ProductService) AddProduct(model models.Product) error {
	err := svc.repository.AddProduct(model)
	if err != nil {
		return err
	}
	return nil
}
func (svc *ProductService) GetProductById(id string) (*models.Product, error) {
	product, err := svc.repository.GetProdcutByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (svc *ProductService) DeleteProduct(id string) error {
	if err := svc.repository.DeleteProduct(id); err != nil {
		return err
	}
	return nil
}
func (svc *ProductService) UpdateProduct(id string, product models.Product) error {
	if err := svc.repository.UpdateProduct(id , product); err != nil {
		return err
	}
	return nil
}