package Routes

import (
	"net/http"

	"github.com/DonShanilka/tvSeries-service/internal/Handler"
)

func RegisterTvSeriesRoutes(mux *http.ServeMux, handler *Handler.TvSeriesHandler) {
	mux.HandleFunc("/api/series/createSeries", handler.CreateTvSeries)
	mux.HandleFunc("/api/series/updateSeries", handler.UpdateTvSeries)
	mux.HandleFunc("/api/series/deleteSeries", handler.DeleteTvSeries)
	mux.HandleFunc("/api/series/getByIdSeries", handler.GetTvSeriesByID)
	mux.HandleFunc("/api/series/getAllSeries", handler.GetAllTvSeries)

}

func RegisterEpisodeRoutes(mux *http.ServeMux, h *Handler.EpisodeHandler) {
	// Episode routes
	mux.HandleFunc("/api/episode/createEpisode", h.CreateEpisode)
	mux.HandleFunc("/api/episode/updateEpisode", h.UpdateEpisode)
	mux.HandleFunc("/api/episode/deleteEpisode", h.DeleteEpisode)
	mux.HandleFunc("/api/episode/getAllEpisode", h.GetAllEpisodes)
	mux.HandleFunc("/api/episode/getEpisodeById", h.GetEpisodeByID)
	mux.HandleFunc("/api/episodes/bySeriesId", h.GetEpisodesBySeriesID)

}
