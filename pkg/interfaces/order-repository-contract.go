package interfaces

import "ecommerce/pkg/models"

type OrderRepositoryContract interface {
	AddOrder(order *models.Orders) error
}
