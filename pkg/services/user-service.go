package services

import (
	"ecommerce/pkg/interfaces"
	"ecommerce/pkg/models"
	"ecommerce/pkg/repository"
	"ecommerce/pkg/utils"
	"errors"
)

type UserService struct {
	repository interfaces.UserRepositoryContract
}

func NewUserService() *UserService {
	repository := repository.NewPostgres()
	return &UserService{
		repository: repository,
	}
}
func (r *UserService) Login(email string, password string) (string, error) {
	user, err := r.repository.Login(email)

	if err != nil {
		return "", err
	}
	isPasswordMatched := utils.CheckPasswordHash(password, user.Password)

	if !isPasswordMatched {
		return "", errors.New("email or password is not correct")
	}

	token , err := utils.TokenGenerator(*user)

	if err != nil {
		return "" , err
	}

	return token, nil
}
func (r *UserService) Register(user models.User) error {

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := r.repository.Register(user); err != nil {
		return err
	}
	return nil
}
