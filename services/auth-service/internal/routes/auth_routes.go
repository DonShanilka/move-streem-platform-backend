package routes

import (
	"github.com/DonShanilka/auth-service/internal/config"
	"github.com/DonShanilka/auth-service/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthRoutes(app *fiber.App, db *mongo.Database, cfg *config.Config) {

	// Initialize the handler with EXISTING DB + Config
	authHandler := handlers.InitAuthHandler(db, cfg)

	// Group /api/auth Routes
	authGroup := app.Group("/api/auth")

	authGroup.Post("/register", authHandler.Register)
	// authGroup.Post("/login", authHandler.Login)
}
