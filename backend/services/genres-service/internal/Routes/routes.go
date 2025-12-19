package Routes

import (
	"net/http"

	"github.com/DonShanilka/genres-service/internal/Handler"
)

func RegisterGenreRoutes(mux *http.ServeMux, handler *Handler.GenreHandler) {
	mux.HandleFunc("/api/genre/creatGenres", handler.CreateGenre)
	mux.HandleFunc("/api/genre/getAllGenre", handler.GetAllGenres)
	mux.HandleFunc("/api/genre/updateGenre", handler.UpdateGenre)
	mux.HandleFunc("/api/genre/deleteGenre", handler.DeleteGenre)
}
