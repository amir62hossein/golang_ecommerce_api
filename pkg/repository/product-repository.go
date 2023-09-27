package repository

import (
	"ecommerce/pkg/models"
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func (p *Postgres) AddProduct(product models.Product) error {

	if err := p.db.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *Postgres) GetProdcutByID(id string) (*models.Product, error) {
	var product models.Product
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	if err := p.db.Where("id = ?", uuid).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
func (p *Postgres) DeleteProduct(id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	if err := p.db.Where("id = ?", uuid).Delete(&models.Product{}).Unscoped().Error; err != nil {
		return err
	}
	return nil
}
func (p *Postgres) UpdateProduct(id string, product models.Product) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	existingProduct := models.Product{}
	if err := p.db.First(&existingProduct, uuid).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("product not found")
		}
		return err
	}
	existingProduct.Name = product.Name
	existingProduct.Description = product.Description
	existingProduct.Price = product.Price
	existingProduct.CategoryID = product.CategoryID

	// Save the updated product to the database
	if err := p.db.Save(&existingProduct).Error; err != nil {
		return err
	}

	return nil
}
