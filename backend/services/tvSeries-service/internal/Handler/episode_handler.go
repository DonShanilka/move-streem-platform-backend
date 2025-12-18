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

func NewEpisodeHandler(s *Service.EpisodeService) *EpisodeHandler {
	return &EpisodeHandler{Service: s}
}

// POST /api/episodes/upload
func (h *EpisodeHandler) CreateEpisode(w http.ResponseWriter, r *http.Request) {
	// Parse multipart/form-data
	if err := r.ParseMultipartForm(100 << 20); err != nil { // 100MB limit
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("episode")
	if err != nil {
		http.Error(w, "Video file is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get metadata
	ep := Models.Episode{
		SeriesID:      atoiSafe(r.FormValue("series_id")), // get from form or URL
		SeasonNumber:  atoiSafe(r.FormValue("season_number")),
		EpisodeNumber: atoiSafe(r.FormValue("episode_number")),
		Title:         r.FormValue("title"),
		Description:   r.FormValue("description"),
		Duration:      atoiSafe(r.FormValue("duration")),
		ReleaseDate:   r.FormValue("release_date"),
	}

	// Upload to B2 + save in MySQL
	err = h.Service.CreateEpisode(&ep, file, header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ep)
}

func (h *EpisodeHandler) UpdateEpisode(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(100 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("episode")
	if err != nil {
		http.Error(w, "Video file required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ep := Models.Episode{
		ID:            atoiSafe(r.FormValue("id")),
		SeriesID:      atoiSafe(r.FormValue("series_id")),
		SeasonNumber:  atoiSafe(r.FormValue("season_number")),
		EpisodeNumber: atoiSafe(r.FormValue("episode_number")),
		Title:         r.FormValue("title"),
		Description:   r.FormValue("description"),
		Duration:      atoiSafe(r.FormValue("duration")),
		ReleaseDate:   r.FormValue("release_date"),
	}

	if err := h.Service.UpdateEpisode(&ep, file, header.Filename); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ep)
}

func (h *EpisodeHandler) DeleteEpisode(w http.ResponseWriter, r *http.Request) {
	id := atoiSafe(r.URL.Query().Get("id"))

	if err := h.Service.DeleteEpisode(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Episode deleted successfully"))
}
