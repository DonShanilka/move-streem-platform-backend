package main

import (
	"log"

	"github.com/DonShanilka/api-gateway/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Register routes
	RegisterRoutes(app)

	log.Println("API Gateway running on :8080")
	log.Fatal(app.Listen(":8080"))
}

func RegisterRoutes(app *fiber.App) {
	routes.RegisterRoutes(app)
}
