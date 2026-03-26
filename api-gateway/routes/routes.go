package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// Forward Auth requests
	api.All("/auth/*", proxyAuthService)

	// Forward User requests
	api.All("/users/*", proxyUserService)

	// Forward Movie requests
	api.All("/movies/*", proxyMovieService)

	// Forward Admin requests
	api.All("/admin/*", proxyAdminService)

	// Forward Genres requests
	api.All("/genres/*", proxyGenresService)

	// Forward Payment requests
	api.All("/payment/*", proxyPaymentService)

	// Forward TV Series requests
	api.All("/tvseries/*", proxyTVSeriesService)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// ReverseProxy forwards request to target microservice
func ReverseProxy(target string, c *fiber.Ctx) error {
	req := c.Request()

	// Keep original path after /api/v1/
	req.SetRequestURI(string(c.Request().RequestURI())[len("/api/v1"):])
	req.SetHost(target)

	client := fasthttp.Client{}
	if err := client.Do(req, c.Response()); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Service unreachable"})
	}
	return nil
}

func proxyAuthService(c *fiber.Ctx) error {
	target := getEnv("AUTH_SERVICE_URL", "localhost:8081")
	return ReverseProxy(target, c)
}

func proxyUserService(c *fiber.Ctx) error {
	target := getEnv("USER_SERVICE_URL", "localhost:8082")
	return ReverseProxy(target, c)
}

func proxyAdminService(c *fiber.Ctx) error {
	target := getEnv("ADMIN_SERVICE_URL", "localhost:8083")
	return ReverseProxy(target, c)
}

func proxyMovieService(c *fiber.Ctx) error {
	target := getEnv("MOVIE_SERVICE_URL", "localhost:8080")
	return ReverseProxy(target, c)
}

func proxyGenresService(c *fiber.Ctx) error {
	target := getEnv("GENRES_SERVICE_URL", "localhost:8084")
	return ReverseProxy(target, c)
}

func proxyPaymentService(c *fiber.Ctx) error {
	target := getEnv("PAYMENT_SERVICE_URL", "localhost:8085")
	return ReverseProxy(target, c)
}

func proxyTVSeriesService(c *fiber.Ctx) error {
	target := getEnv("TVSERIES_SERVICE_URL", "localhost:8086")
	return ReverseProxy(target, c)
}
