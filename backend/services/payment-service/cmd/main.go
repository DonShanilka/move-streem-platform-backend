package main

import (
	"log"
	"net/http"

	//"net/http"

	"backend/payment-service/Middleware"
	"backend/payment-service/internal/Handler"
	"backend/payment-service/internal/Repository"
	"backend/payment-service/internal/Routes"
	"backend/payment-service/internal/Service"
	"backend/payment-service/internal/db"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	planRepo := Repository.NewPlanRepostry(database)
	planService := Service.NewPlanService(planRepo)
	planHandler := Handler.NewGenreHandler(planService)

	subsRepo := Repository.NewSubsRepostry(database)
	subsService := Service.NewSubsService(subsRepo)
	subsHandler := Handler.NewSubsHandler(subsService)

	mux := http.NewServeMux()
	Routes.RegisterPlanRoutes(mux, planHandler)
	Routes.RegisterSubsRoutes(mux, subsHandler)

	log.Println("Plan & Subs Service running on :8080 🚀")
	err = http.ListenAndServe(":8080", Middleware.CorsMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}
