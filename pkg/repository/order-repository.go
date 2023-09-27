package repository

import "ecommerce/pkg/models"

func (p *Postgres) AddOrder(order *models.Orders) error {
	if err := p.db.Create(&order).Error; err != nil {
		return err
	}
	return nil
}
