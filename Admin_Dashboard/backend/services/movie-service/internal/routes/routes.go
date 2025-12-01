package routes

import (
    "net/http"
    "github.com/DonShanilka/movie-service/internal/handlers"
)

// SetupRoutes registers all routes and returns a handler
func SetupRoutes(movieHandler *handlers.MovieHandler) http.Handler {
    mux := http.NewServeMux()

    // Movie routes
    mux.HandleFunc("/upload", movieHandler.UploadMovie)
    mux.Handle("/videos/", http.StripPrefix("/videos/", http.FileServer(http.Dir("./videos"))))
    mux.HandleFunc("/movies", movieHandler.ListMovies)

    return mux
}
