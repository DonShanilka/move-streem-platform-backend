package main

import (
	"log"
	"net/http"

	"github.com/DonShanilka/genres-service/Middleware"
	"github.com/DonShanilka/genres-service/internal/Handler"
	"github.com/DonShanilka/genres-service/internal/Repository"
	"github.com/DonShanilka/genres-service/internal/Routes"
	"github.com/DonShanilka/genres-service/internal/Service"
	"github.com/DonShanilka/genres-service/internal/db"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	genreRepo := Repository.NewGenerRepostry(database)
	genreService := Service.NewGenreService(genreRepo)
	genreHandler := Handler.NewGenreHandler(genreService)

	mux := http.NewServeMux()
	Routes.RegisterGenreRoutes(mux, genreHandler)

	log.Println("Genres Service running on :8080 🚀")
	err = http.ListenAndServe(":8080", Middleware.CorsMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}
