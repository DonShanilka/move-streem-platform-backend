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
	// Initialize MongoDB Atlas
	database, err := db.InitMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB Atlas ❌:", err)
	}

	// =========================
	// Repositories
	// =========================
	movieRepo := Repository.NewMovieRepository(database.Collection("movies"))
	tvSeriesRepo := Repository.NewTvSeriesRepository(database)

	// =========================
	// Handlers
	// =========================
	// Movies
	movieHandler := Handler.NewMovieHandler(movieRepo)
	// TV Series
	tvSeriesHandler := Handler.NewTvSeriesHandler(tvSeriesRepo)

	// =========================
	// HTTP Routes
	// =========================
	mux := http.NewServeMux()
	Routes.RegisterMovieRoutes(mux, movieHandler)
	Routes.RegisterTvSeriesRoutes(mux, tvSeriesHandler)

	// =========================
	// CORS Middleware
	// =========================
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		mux.ServeHTTP(w, r)
	})

	log.Println("Movie & TV Series Service running on :8080 🚀")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
