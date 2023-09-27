package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Admin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "missing token"})
		}
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "invalid token"})
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			role, ok := claims["ROLE"].(string)
			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "invalid token"})
			}

			if role == "ADMIN" {
				return c.Next()
			}

		}
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "invalid token"})
	}
}
