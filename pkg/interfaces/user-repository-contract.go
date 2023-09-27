package interfaces

import "ecommerce/pkg/models"

type UserRepositoryContract interface {
	Login(email string) (*models.User, error)
	Register(models.User) error
}
