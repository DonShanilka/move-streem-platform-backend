package routes

import (
    "net/http"
    "github.com/DonShanilka/movie-service/internal/handlers"
)

func RegisterMovieRoutes(mux *http.ServeMux, movieHandler *handlers.MovieHandler) {

    // API Routes
    mux.HandleFunc("/api/movies/upload", movieHandler.UploadMovie)
    mux.HandleFunc("/api/movies", movieHandler.ListMovies)

    // Correct video serving
    mux.Handle("/videos/",
        http.StripPrefix("/videos/",
            http.FileServer(http.Dir("./videos")),
        ),
    )
}

