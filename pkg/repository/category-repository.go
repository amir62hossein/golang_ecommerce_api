package repository

import "ecommerce/pkg/models"

func (p *Postgres) AddCategory(category models.Category) error {
	if err := p.db.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func (p *Postgres) GetAllCategory() (*[]models.Category, error) {
	var categories []models.Category

	if err := p.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}
