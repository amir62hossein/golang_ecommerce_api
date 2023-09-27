package repository

import (
	"ecommerce/pkg/models"

	"fmt"
)

func (p *Postgres) Register(user models.User) error {
	if err := p.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (p *Postgres) Login(email string) (*models.User, error) {
	var fundedUser models.User
	if err := p.db.Where("email = ? ", email).First(&fundedUser).Error; err != nil {
		return nil, fmt.Errorf("user with email %s not found", email)
	}
	return &fundedUser, nil

}
