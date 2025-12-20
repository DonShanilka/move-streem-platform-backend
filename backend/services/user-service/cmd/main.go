package main

import (
	"log"
	"net/http"

	"github.com/DonShanilka/user-service/Middleware"
	"github.com/DonShanilka/user-service/internal/Handler"
	"github.com/DonShanilka/user-service/internal/Repository"
	"github.com/DonShanilka/user-service/internal/Routes"
	"github.com/DonShanilka/user-service/internal/Service"
	"github.com/DonShanilka/user-service/internal/db"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := Repository.NewUserRepository(database)
	userService := Service.NewUserService(userRepo)
	userHandler := Handler.NewUserHandler(userService)

	mux := http.NewServeMux()
	Routes.RegisterAdminRoutes(mux, userHandler)

	log.Println("User Service running on :8080 🚀")
	err = http.ListenAndServe(":8080", Middleware.CorsMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}
