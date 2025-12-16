package routes

import (
	"net/http"

	"github.com/DonShanilka/movie-service/internal/handlers"
)

func RegisterMovieRoutes(mux *http.ServeMux, h *handlers.MovieHandler) {
	mux.HandleFunc("/api/movies/getAllMovies", h.ListMovies)
	mux.HandleFunc("/api/movies/upload", h.UploadMovie)
	mux.HandleFunc("/api/movies/stream", h.StreamMovie)
	mux.HandleFunc("/api/movies/update", h.UpdateMovie)
}

func RegisterSeriesRoutes(mux *http.ServeMux, h *handlers.SeriesHandler) {
	mux.HandleFunc("/api/series", h.UpdaloadSeries)
	// mux.HandleFunc("/api/series/upload", h.UploadSeries)
	// mux.HandleFunc("/api/series/stream", h.StreamSeries)
}

func RegisterEpisodeRoutes(mux *http.ServeMux, h *handlers.EpisodeHandler) {
	// mux.HandleFunc("/api/episodes", h.ListEpisodes)
	mux.HandleFunc("/api/episodes/upload", h.UploadEpisode)
	// mux.HandleFunc("/api/episodes/stream", h.StreamEpisode)
}
