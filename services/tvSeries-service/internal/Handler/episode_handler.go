package Handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func (h *EpisodeHandler) GetAllEpisodes(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Service.GetAllEpisodes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(movies)
}

// GET /api/episodes/{id}
func (h *EpisodeHandler) GetEpisodeByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "episode id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid episode id", http.StatusBadRequest)
		return
	}

	episode, err := h.Service.GetEpisodeByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(episode)
}

func (h *EpisodeHandler) GetEpisodesBySeriesID(w http.ResponseWriter, r *http.Request) {
	seriesIDStr := r.URL.Query().Get("seriesId")
	if seriesIDStr == "" {
		http.Error(w, "seriesId is required", http.StatusBadRequest)
		return
	}

	seriesID, err := strconv.Atoi(seriesIDStr)
	if err != nil {
		http.Error(w, "invalid seriesId", http.StatusBadRequest)
		return
	}

	episodes, err := h.Service.GetEpisodesBySeriesID(seriesID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(episodes)
}
