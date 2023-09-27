package interfaces

import "ecommerce/pkg/models"

type OrderServiceContract interface {
	AddOrder(order *models.Orders) error
}
