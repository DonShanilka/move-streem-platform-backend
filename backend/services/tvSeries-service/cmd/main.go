package main

import (
	"log"
	"net/http"

	"github.com/DonShanilka/movie-service/internal/Handler"
	"github.com/DonShanilka/movie-service/internal/Repository"
	"github.com/DonShanilka/movie-service/internal/Routes"
	"github.com/DonShanilka/movie-service/internal/db"
)

func main() {
	// Connect to MongoDB Atlas
	database, err := db.InitMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB Atlas ❌:", err)
	}

	// Create mux and repository/handler
	mux := http.NewServeMux()
	tvSeriesRepo := Repository.NewTvSeriesRepository(database)
	tvSeriesHandler := Handler.NewTvSeriesHandler(tvSeriesRepo)
	Routes.RegisterTvSeriesRoutes(mux, tvSeriesHandler) // register route

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
