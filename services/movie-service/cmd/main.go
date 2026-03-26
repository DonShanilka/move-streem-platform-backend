package main

import (
	"log"
	"net/http"

	"github.com/DonShanilka/movie-service/internal/Handler"
	"github.com/DonShanilka/movie-service/internal/Repository"
	"github.com/DonShanilka/movie-service/internal/Routes"
	"github.com/DonShanilka/movie-service/internal/Service"
	"github.com/DonShanilka/movie-service/internal/db"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Backblaze B2 Config
	b2KeyID := "f9f45a6c989e"
	b2AppKey := "00563942506fbf1481548bd202ea51e42ec0ce19b7"
	b2BucketName := "movieStream"

	movieRepo, err := Repository.NewMovieRepository(database, b2KeyID, b2AppKey, b2BucketName)
	if err != nil {
		log.Fatal("Failed to create Movie Repository", err)
	}
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
