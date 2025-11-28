package main

import (
	//"log"

	"github.com/gofiber/fiber/v2"
	"github.com/DonShanilka/auth-service/internal/routes"
)

func main() {
    app := fiber.New()
    routes.SetupAuthRoutes(app)
    app.Listen(":9002")
}
