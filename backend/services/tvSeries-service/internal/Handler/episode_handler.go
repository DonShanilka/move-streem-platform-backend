package Handler

import (
	"encoding/json"
	"net/http"

	"github.com/DonShanilka/tvSeries-service/internal/Models"
	"github.com/DonShanilka/tvSeries-service/internal/Service"
)

type EpisodeHandler struct {
	Service *Service.EpisodeService
}

func (h *EpisodeHandler) CreateEpisode(w http.ResponseWriter, r *http.Request) {
	var episode Models.Episode

	if err := json.NewDecoder(r.Body).Decode(&episode); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateEpisode(&episode); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(episode)
}
