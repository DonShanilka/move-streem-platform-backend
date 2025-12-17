package Routes

import (
	"net/http"

	"github.com/DonShanilka/movie-service/internal/Handler"
)

func RegisterTvSeriesRoutes(mux *http.ServeMux, h *Handler.TvSeriesHandler) {
	mux.HandleFunc("/api/series/createSeries", h.CreateSeries) // POST
}
