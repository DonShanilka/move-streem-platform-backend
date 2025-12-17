package Handler

import (
	"encoding/json"
	"net/http"

	"github.com/DonShanilka/tvSeries-service/internal/Models"
	"github.com/DonShanilka/tvSeries-service/internal/Service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TvSeriesHandler struct {
	Service *Service.TvSeriesService
}

func NewTvSeriesHandler(
	service *Service.TvSeriesService,
) *TvSeriesHandler {
	return &TvSeriesHandler{Service: service}
}

// ---------------- CREATE SERIES ----------------
func (h *TvSeriesHandler) CreateSeries(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	var series Models.Series
	if err := json.NewDecoder(r.Body).Decode(&series); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	id, err := h.Service.CreateSeries(&series)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Series created",
		"id":      id.Hex(),
	})
}

// ---------------- ADD SEASON ----------------
func (h *TvSeriesHandler) AddSeason(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	seriesIDHex := r.URL.Query().Get("seriesId")
	if seriesIDHex == "" {
		http.Error(w, "seriesId required", http.StatusBadRequest)
		return
	}

	seriesID, err := primitive.ObjectIDFromHex(seriesIDHex)
	if err != nil {
		http.Error(w, "Invalid seriesId", http.StatusBadRequest)
		return
	}

	var season Models.Season
	if err := json.NewDecoder(r.Body).Decode(&season); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.Service.AddSeason(seriesID, season); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Season added successfully",
	})
}
