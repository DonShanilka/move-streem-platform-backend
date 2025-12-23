package Routes

import (
	"net/http"

	"github.com/DonShanilka/tvSeries-service/internal/Handler"
)

func RegisterTvSeriesRoutes(mux *http.ServeMux, handler *Handler.TvSeriesHandler) {
	mux.HandleFunc("/tv-series/create", handler.CreateTvSeries)
	mux.HandleFunc("/tv-series/update", handler.UpdateTvSeries)
	mux.HandleFunc("/tv-series/delete", handler.DeleteTvSeries)
	mux.HandleFunc("/tv-series/get", handler.GetTvSeriesByID)
	mux.HandleFunc("/tv-series/list", handler.GetAllTvSeries)

}

func RegisterEpisodeRoutes(mux *http.ServeMux, h *Handler.EpisodeHandler) {
	// Episode routes
	mux.HandleFunc("/api/episodes/create", h.CreateEpisode)
	mux.HandleFunc("/api/episodes/update", h.UpdateEpisode)
	mux.HandleFunc("/api/episodes/delete", h.DeleteEpisode)
	mux.HandleFunc("/api/episodes/getAllEpisode", h.GetAllEpisodes)
	mux.HandleFunc("/api/episodes/getEpisodeById", h.GetEpisodeByID)
}
