package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// Forward Auth requests
	api.All("/auth/*", proxyAuthService)
}

// ReverseProxy forwards request to target microservice
func ReverseProxy(target string, c *fiber.Ctx) error {
	req := c.Request()

	// Keep original path after /auth/
	req.SetRequestURI(string(c.Request().RequestURI())[len("/api/v1"):])
	req.SetHost(target)

	client := fasthttp.Client{}
	if err := client.Do(req, c.Response()); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Service unreachable"})
	}
	return nil
}

func proxyAuthService(c *fiber.Ctx) error {
	return ReverseProxy("localhost:9002", c) // Auth Service port
}
