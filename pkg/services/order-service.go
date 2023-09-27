package services

import (
	"ecommerce/pkg/interfaces"
	"ecommerce/pkg/models"
	"ecommerce/pkg/repository"
)

type OrderService struct {
	repo interfaces.OrderRepositoryContract
}

func NewOrderService() *OrderService {
	repo := repository.NewPostgres()
	return &OrderService{
		repo: repo,
	}
}

func (r *OrderService) AddOrder(order *models.Orders) error {
	if err := r.repo.AddOrder(order); err != nil {
		return err
	}
	return nil
}
