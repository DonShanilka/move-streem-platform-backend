package routes

import (
    "net/http"
    "github.com/DonShanilka/movie-service/internal/handlers"
)

func RegisterMovieRoutes(mux *http.ServeMux, h *handlers.MovieHandler) {
    mux.HandleFunc("/api/movies", h.ListMovies)
    mux.HandleFunc("/api/movies/upload", h.UploadMovie)
    mux.HandleFunc("/api/movies/stream", h.StreamMovie)
}
