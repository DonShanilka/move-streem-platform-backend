package Handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/DonShanilka/movie-service/internal/Models"
	"github.com/DonShanilka/movie-service/internal/Repository"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
)

type TvSeriesHandler struct {
	Repo *Repository.TvSeriesRepository
}

func NewTvSeriesHandler(repo *Repository.TvSeriesRepository) *TvSeriesHandler {
	return &TvSeriesHandler{Repo: repo}
}

// Create a new TV series
func (h *TvSeriesHandler) CreateSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var series Models.Series

	// Parse JSON body
	if err := json.NewDecoder(r.Body).Decode(&series); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Optional: Set default values
	series.SeasonCount = len(series.Seasons)
	series.Seasons = nil // We store episodes separately

	// Insert into MongoDB
	id, err := h.Repo.CreateSeries(&series)
	if err != nil {
		log.Println("Error creating series:", err)
		http.Error(w, "Failed to create series", http.StatusInternalServerError)
		return
	}

	// Return success response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "TV series created successfully",
		"id":      id.Hex(),
		"created": time.Now().Format(time.RFC3339),
	})
}
