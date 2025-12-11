package main

import (
	"log"
	"net/http"

	"github.com/DonShanilka/movie-service/internal/handlers"
	"github.com/DonShanilka/movie-service/internal/db"
	"github.com/DonShanilka/movie-service/internal/routes"
	"github.com/DonShanilka/movie-service/internal/service"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	movieService := services.NewMovieService(db)
	movieHandler := handlers.NewMovieHandler(movieService)

	seriesService := services.NewSeriesService(db)
	seriesHandler := handlers.NewSeriesHandler(seriesService)

	episiodeService := services.NewEpisodeService(db)
	episodeHandler := handlers.NewEpisodeHandler(episiodeService)

	mux := http.NewServeMux()
	routes.RegisterMovieRoutes(mux, movieHandler)
	routes.RegisterSeriesRoutes(mux, seriesHandler)
	routes.RegisterEpisodeRoutes(mux, episodeHandler)

	// Global CORS middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			return
		}

		mux.ServeHTTP(w, r)
	})

	log.Println("Server running at http://localhost:8080 ✅")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
