package Routes

import (
	"net/http"

	"github.com/DonShanilka/admin-service/internal/Handler"
)

func RegisterAdminRoutes(mux *http.ServeMux, handler *Handler.AdminHandler) {
	mux.HandleFunc("/api/genre/creatGenres", handler.CreateAdmin)
	mux.HandleFunc("/api/genre/getAllGenre", handler.)
	mux.HandleFunc("/api/genre/updateGenre", handler.UpdateGenre)
	mux.HandleFunc("/api/genre/deleteGenre", handler.DeleteGenre)
}
