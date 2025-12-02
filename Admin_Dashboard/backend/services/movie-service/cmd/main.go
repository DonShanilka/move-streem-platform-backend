package main

import (
    "log"
    "net/http"

    "github.com/DonShanilka/movie-service/internal/handlers"
    "github.com/DonShanilka/movie-service/internal/repository"
    "github.com/DonShanilka/movie-service/internal/routes"
    "github.com/DonShanilka/movie-service/internal/services"
)

func main() {
    db, err := repository.InitDB()
    if err != nil {
        log.Fatal(err)
    }

    movieService := services.NewMovieService(db)
    movieHandler := handlers.NewMovieHandler(movieService)

    mux := http.NewServeMux()
    routes.RegisterMovieRoutes(mux, movieHandler)

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
