package main

import (
	"log"
	"net/http"

	"github.com/DonShanilka/tvSeries-service/internal/Handler"
	"github.com/DonShanilka/tvSeries-service/internal/Repository"
	"github.com/DonShanilka/tvSeries-service/internal/Routes"
	"github.com/DonShanilka/tvSeries-service/internal/Service"
	"github.com/DonShanilka/tvSeries-service/internal/db"
)

func main() {
	// Connect to MongoDB Atlas
	database, err := db.InitMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB Atlas ❌:", err)
	}

	// Repository
	tvSeriesRepo := Repository.NewTvSeriesRepository(database)

	// Service
	tvSeriesService := Service.NewTvSeriesService(tvSeriesRepo)

	// Handler
	tvSeriesHandler := Handler.NewTvSeriesHandler(tvSeriesService)

	mux := http.NewServeMux()
	Routes.RegisterTvSeriesRoutes(mux, tvSeriesHandler)

	// Start server with CORS wrapper
	log.Println("TV Series Service running on :8080 🚀")
	err = http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Route the request
		mux.ServeHTTP(w, r)
	}))
	if err != nil {
		log.Fatal(err)
	}
}
