// auth-service/internal/routes/routes.go
package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/DonShanilka/auth-service/internal/handlers"
)

func SetupAuthRoutes(app *fiber.App) {
    api := app.Group("/auth")
    api.Post("/register", handlers.Register)
    api.Post("/login", handlers.Login)
}
