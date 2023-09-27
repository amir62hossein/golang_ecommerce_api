package main

import (
	"ecommerce/pkg/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic("Failed To Load .env Files")
	}

	app := fiber.New()
	
	handlers.SetupUserRoutes(app)
	handlers.SetupProductRoutes(app)
	handlers.SetupCategoryRoutes(app)
	handlers.SetupCommentRoutes(app)
	handlers.SetupOrderRoutes(app)

	log.Fatal(app.Listen(":8080"))

}
