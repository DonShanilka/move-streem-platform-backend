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
func (h *EpisodeHandler) UploadEpisode(w http.ResponseWriter, r *http.Request) {
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
	err = h.Service.UploadEpisode(&ep, file, header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ep)
}

//func atoiSafe(value string) int {
//	i, _ := strconv.Atoi(value)
//	return i
//}
