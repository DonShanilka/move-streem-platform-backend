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

	title := r.FormValue("title")
	description := r.FormValue("description")
	duration := r.FormValue("duration")
	thumbnailURL := r.FormValue("thumbnailURL")
	// episodeURL := r.FormValue("episodeURL")
	releaseDate := r.FormValue("releaseDate")

	// ========== SAVE EPISODE TO LOCAL DISK ==========
	episodeFile, episodeHeader, err := r.FormFile("episode")
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

	// =========== create model =================
	episode := models.Episode {
		Title:        title,
		Description:  description,
		Duration:     atoiSafe(duration),
		ThumbnailURL: thumbnailURL,
		ReleaseDate:  releaseDate,
		EpisodeURL:   episodeHeader.Filename,
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
