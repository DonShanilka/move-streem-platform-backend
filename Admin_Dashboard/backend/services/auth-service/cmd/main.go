package main

import (
	"log"

	"github.com/DonShanilka/auth-service/internal/routes"
	"github.com/DonShanilka/auth-service/internal/config"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	routes.AuthRoutes(app, cfg)

	log.Println("Auth Service running on port:", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
