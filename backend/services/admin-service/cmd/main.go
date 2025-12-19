package main

import (
	"log"
	"net/http"

	"github.com/DonShanilka/admin-service/Middleware"
	"github.com/DonShanilka/admin-service/internal/Handler"
	"github.com/DonShanilka/admin-service/internal/Repository"
	"github.com/DonShanilka/admin-service/internal/Routes"
	"github.com/DonShanilka/admin-service/internal/Service"
	"github.com/DonShanilka/admin-service/internal/db"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	adminRepo := Repository.NewAdminRepository(database)
	adminService := Service.NewAdminService(adminRepo)
	adminHandler := Handler.NewAdminHandler(adminService)

	mux := http.NewServeMux()
	Routes.RegisterAdminRoutes(mux, adminHandler)

	log.Println("Genres Service running on :8080 🚀")
	err = http.ListenAndServe(":8080", Middleware.CorsMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}
