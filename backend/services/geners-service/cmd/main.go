package main

import (
	"log"
	"net/http"

	"github.com/DonShanilka/geners-service/internal/Handler"
	"github.com/DonShanilka/geners-service/internal/Repository"
	"github.com/DonShanilka/geners-service/internal/Routes"
	"github.com/DonShanilka/geners-service/internal/Service"
	"github.com/DonShanilka/geners-service/internal/db"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	movieRepo := Repository.NewMovieRepository(database)
	movieService := services.NewMovieService(movieRepo)
	movieHandler := Handler.NewMovieHandler(movieService)

	mux := http.NewServeMux()
	Routes.RegisterMovieRoutes(mux, movieHandler)

	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if request.Method == http.MethodOptions {
			return
		}
		mux.ServeHTTP(writer, request)
	})

	log.Println("Movie Service running on :8080 🚀")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
