package routes

import (
    "github.com/DonShanilka/auth-service/internal/handlers"
    "github.com/gofiber/fiber/v2"
    "github.com/DonShanilka/auth-service/internal/config"

    "go.mongodb.org/mongo-driver/mongo"
)

func AuthRoutes(app *fiber.App, cfg *config.Config) {
    client, _ := mongo.Connect(nil)
    db := client.Database(cfg.Database)
    authHandler := handlers.InitAuthHandler(db, cfg)

    authGroup := app.Group("/api/auth")
    authGroup.Post("/register", authHandler.Register)
    // authGroup.Post("/login", authHandler.Login)
}