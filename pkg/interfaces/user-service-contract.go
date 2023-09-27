package interfaces

import "ecommerce/pkg/models"

type UserServiceContract interface {
	Login(email string , password string) (string, error)
	Register(models.User) error
}
