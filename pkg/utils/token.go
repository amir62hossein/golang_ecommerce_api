package utils

import (
	"ecommerce/pkg/models"
	"os"
	"github.com/golang-jwt/jwt/v5"
)

func TokenGenerator(user models.User) (string, error) {

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":   user.ID,
		"ROLE": user.Role,
	})
	tokenString, err := t.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
