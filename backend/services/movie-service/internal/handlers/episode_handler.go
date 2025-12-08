package handlers

import (
	"io"
	"net/http"
	"os"

	"github.com/DonShanilka/movie-service/internal/models"
	"github.com/DonShanilka/movie-service/internal/service"
)

type EpisodeHandler struct {
	Episode *services.EpisodeService
}

func NewEpisodeHandler(episode *services.EpisodeService) *EpisodeHandler {
	return &EpisodeHandler{Episode: episode}
} 

func (h *EpisodeHandler) UploadEpisode(w http.ResponseWriter, r * http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(110 << 30)

	seriesID := r.FormValue("SeriesID")
	seasonNumber := r.FormValue("SeasonNumber")
	episodeNumber := r.FormValue("EpisodeNumber")
	title := r.FormValue("Title")
	description := r.FormValue("Description")
	duration := r.FormValue("Duration")
	thumbnailURL := r.FormValue("ThumbnailURL")
	releaseDate := r.FormValue("ReleaseDate")

	// ========== SAVE EPISODE TO LOCAL DISK ==========
	episodeFile, episodeHeader, err := r.FormFile("EpisodeURL")
	if err != nil {
		http.Error(w, "Episode File missing: "+err.Error(), 400)
		return
	}
	defer episodeFile.Close()

	episodePath := "./episodes/" + episodeHeader.Filename

	os.MkdirAll("./episodes", 0755)
	f, err := os.Create(episodePath)
	if err != nil {
		http.Error(w, "Failed to create episode file: "+err.Error(), 500)
		return
	}
	io.Copy(f, episodeFile)
	f.Close()

	episode := models.Episode{
    SeriesID:      atoiSafe(seriesID),
    SeasonNumber:  atoiSafe(seasonNumber),
    EpisodeNumber: atoiSafe(episodeNumber),
    Title:         title,
    Description:   description,
    Duration:      atoiSafe(duration),
    ThumbnailURL:  thumbnailURL,
    EpisodeURL:    episodeHeader.Filename,
    ReleaseDate:   releaseDate,
}

	if err := h.Episode.SaveEpisode(episode); err != nil {
		http.Error(w, "DB Error: " + err.Error(), 500)
		return
	}

	jsonResponse(w, map[string]interface{} {
		"message": "Episode uploaded successfully",
		"episode_local": episodePath,
	})
}
