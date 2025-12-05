package handlers

import (
	// "crypto/des"
	// "encoding/json"
	"io"
	"net/http"
	// "os"
	// "strconv"

	"github.com/DonShanilka/movie-service/internal/models"
	"github.com/DonShanilka/movie-service/internal/service"
)

type SeriesHandler struct {
	Service *services.SeriesService
}

func NewSeriesHandler(series *services.SeriesService) *SeriesHandler {
	return &SeriesHandler{Service: series}
}

func (h *SeriesHandler) UpdaloadSeries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "MEthod Not Allowd ", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(110 << 30)

	title := r.FormValue("title")
	description := r.FormValue("description")
	releasYear := r.FormValue("ReleaseYear")
	language := r.FormValue("laguage")
	seasonCount := r.FormValue("seasonCount")
	thumbnailURL := r.FormValue("thumbnailURL")
	
	// ========== READ BANNER (BLOB) ==========
	bannerFile, _, _ := r.FormFile("banner")
	var banner []byte
	if bannerFile != nil {
		banner, _ = io.ReadAll(bannerFile)
		bannerFile.Close()
	}

	// =========== create model =================
	services := models.Series {
		Title:        title,
		Description:  description,
		ReleaseYear:  atoiSafe(releasYear),
		Language: 	 language,
		SeasonCount:  atoiSafe(seasonCount),
		ThumbnailURL: thumbnailURL,
		Banner: banner,
	}

	if err := h.Service.SaveSeries(services); err != nil {
		http.Error(w, "Failed to save series" + err.Error(), 500)
		return
	}

	jsonResponse(w, map[string]interface{}{
		"message": "Series uploaded successfully",
	})

}