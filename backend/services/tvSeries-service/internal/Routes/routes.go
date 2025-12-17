package Routes

import (
	"net/http"

	"github.com/DonShanilka/movie-service/internal/Handler"
)

// RegisterTvSeriesRoutes registers all TV series-related routes
func RegisterTvSeriesRoutes(mux *http.ServeMux, h *Handler.TvSeriesHandler) {
	// Create a new series
	mux.HandleFunc("/api/series", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.CreateSeries(w, r)
		case http.MethodGet:
			//h.GetAllSeries(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Update a series
	//mux.HandleFunc("/api/series/update", func(w http.ResponseWriter, r *http.Request) {
	//	if r.Method == http.MethodPut {
	//		h.UpdateSeries(w, r)
	//		return
	//	}
	//	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	//})
	//
	//// Delete a series
	//mux.HandleFunc("/api/series/delete", func(w http.ResponseWriter, r *http.Request) {
	//	if r.Method == http.MethodDelete {
	//		h.DeleteSeries(w, r)
	//		return
	//	}
	//	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	//})
	//
	//// Create episode
	//mux.HandleFunc("/api/series/episode", func(w http.ResponseWriter, r *http.Request) {
	//	if r.Method == http.MethodPost {
	//		h.CreateEpisode(w, r)
	//		return
	//	}
	//	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	//})
}
