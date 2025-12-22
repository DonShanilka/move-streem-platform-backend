package main

import (
	"log"

	"github.com/DonShanilka/auth-service/internal/config"
	"github.com/DonShanilka/auth-service/internal/database"
	"github.com/DonShanilka/auth-service/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Starting Auth Service...")

	// Load ENV
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load Config:", err)
	}

	log.Println("PORT:", cfg.Port)
	log.Println("MONGO_URI:", cfg.MongoURI)
	log.Println("MONGO_DB:", cfg.Database)

	// Connect to Mongo (pass BOTH URI and DB name)
	db, err := database.ConnectMongo(cfg.MongoURI, cfg.Database)
	if err != nil {
		log.Fatal("Mongo connection failed:", err)
	}
	log.Println("MongoDB connected!")

	// Setup Fiber
	app := fiber.New()

	// Pass DB + Config into Routes
	routes.AuthRoutes(app, db, cfg)

	// Start server
	log.Println("Auth Service running on port:", cfg.Port)

	err = app.Listen("0.0.0.0:" + cfg.Port)
	if err != nil {
		log.Fatal("Fiber failed:", err)
	}
}
