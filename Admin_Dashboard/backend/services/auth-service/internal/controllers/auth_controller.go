package controllers

import (
	"github.com/DonShanilka/auth-service/internal/models"	
	"github.com/DonShanilka/auth-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON", "details": err.Error()})
	}

	err := c.AuthService.Register(&user)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "User registered successfully"})
}