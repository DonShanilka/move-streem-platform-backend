package main

import (
	"log"
	//"net/http"

	//"github.com/DonShanilka/payment-service/Middleware"
	//"github.com/DonShanilka/payment-service/internal/Handler"
	//"github.com/DonShanilka/payment-service/internal/Repository"
	//"github.com/DonShanilka/payment-service/internal/Routes"
	//"github.com/DonShanilka/payment-service/internal/Service"
	"backend/payment-service/internal/db"
)

func main() {
	_, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	//userRepo := Repository.NewUserRepository(database)
	//userService := Service.NewUserService(userRepo)
	//userHandler := Handler.NewUserHandler(userService)
	//
	//mux := http.NewServeMux()
	//Routes.RegisterUserRoutes(mux, userHandler)

	log.Println("User Service running on :8080 🚀")
	//err = http.ListenAndServe(":8080", Middleware.CorsMiddleware(mux))
	//if err != nil {
	//	log.Fatal(err)
	//}
}
