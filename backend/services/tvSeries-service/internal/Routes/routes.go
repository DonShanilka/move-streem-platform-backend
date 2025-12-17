package Routes

import (
	"net/http"

	"github.com/DonShanilka/tvSeries-service/internal/Handler"
)

func RegisterTvSeriesRoutes(mux *http.ServeMux, h *Handler.TvSeriesHandler) {
	mux.HandleFunc("/api/series/create", h.CreateSeries)  // POST
	mux.HandleFunc("/api/series/add-season", h.AddSeason) // POST
}
