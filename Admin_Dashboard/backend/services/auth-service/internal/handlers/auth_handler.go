package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/DonShanilka/auth-service/internal/service"
)

func Register(c *fiber.Ctx) error {
    req := new(struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    })
    c.BodyParser(req)

    err := services.RegisterUser(req.Name, req.Email, req.Password)
    if err != nil { return c.Status(500).JSON(fiber.Map{"error": err.Error()}) }

    return c.JSON(fiber.Map{"message": "User registered"})
}

func Login(c *fiber.Ctx) error {
    req := new(struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    })
    c.BodyParser(req)

    token, err := services.LoginUser(req.Email, req.Password)
    if err != nil { return c.Status(401).JSON(fiber.Map{"error": "Invalid login"}) }

    return c.JSON(fiber.Map{"token": token})
}
