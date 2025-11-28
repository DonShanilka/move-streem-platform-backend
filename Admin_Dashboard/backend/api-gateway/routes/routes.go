// api-gateway/internal/routes/routes.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// Forward /auth/* to auth-service
	api.Post("/auth/register", proxyAuthService)
	api.Post("/auth/login", proxyAuthService)
}

func ReverseProxy(target string, c *fiber.Ctx) error {
	req := c.Request()
	req.SetRequestURI(string(c.Request().RequestURI()))
	req.SetHost(target)

	client := fasthttp.Client{}
	if err := client.Do(req, c.Response()); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Service unreachable"})
	}
	return nil
}

func proxyAuthService(c *fiber.Ctx) error {
	return ReverseProxy("localhost:9002", c) // port where auth-service runs
}
