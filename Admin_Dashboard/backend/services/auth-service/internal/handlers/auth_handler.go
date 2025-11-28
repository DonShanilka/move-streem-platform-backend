package handlers

import (
	"github.com/DonShanilka/auth-service/internal/config"
	"github.com/DonShanilka/auth-service/internal/controllers"
	"github.com/DonShanilka/auth-service/internal/repository"
	"github.com/DonShanilka/auth-service/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitAuthHandler(db *mongo.Database, cfg *config.Config) *controllers.AuthController {
	userRepo := repository.NewUserRepository(db.Collection("users"))
	authService := services.NewAuthService(*userRepo, cfg.JWTSecret)
	return controllers.NewAuthController(authService)
}
