package Handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/DonShanilka/tvSeries-service/internal/Models"
	"github.com/DonShanilka/tvSeries-service/internal/Service"
)

type TvSeriesHandler struct {
	Service *Service.TvSerriesService
}

func NewTvSeriesHandler(service *Service.TvSerriesService) *TvSeriesHandler {
	return &TvSeriesHandler{Service: service}
}

func (h *TvSeriesHandler) CreateTvSeries(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(100 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tvSeries := Models.Series{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		ReleaseYear: atoiSafe(r.FormValue("releaseYear")),
		SeasonCount: atoiSafe(r.FormValue("seasonCount")),
		Language:    r.FormValue("language"),
	}

	if file, _, err := r.FormFile("banner"); err == nil && file != nil {
		tvSeries.Banner, _ = io.ReadAll(file)
		file.Close()
	}

	if err := h.Service.CreateTvSeries(&tvSeries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(w, http.StatusCreated, map[string]string{
		"status":  "success",
		"message": "TV series created",
	})
}

func (h *TvSeriesHandler) GetAllTvSeries(w http.ResponseWriter, r *http.Request) {
	series, err := h.Service.GetAllTvSeries()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(w, http.StatusOK, series)
}

func (h *TvSeriesHandler) GetTvSeriesByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id := uint(atoiSafe(idStr))

	if id == 0 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	series, err := h.Service.GetTvSeriesByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse(w, http.StatusOK, series)
}

func (h *TvSeriesHandler) UpdateTvSeries(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id := uint(atoiSafe(idStr))

	if id == 0 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(100 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateData := Models.Series{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		ReleaseYear: atoiSafe(r.FormValue("releaseYear")),
		SeasonCount: atoiSafe(r.FormValue("seasonCount")),
		Language:    r.FormValue("language"),
	}

	if file, _, err := r.FormFile("banner"); err == nil && file != nil {
		updateData.Banner, _ = io.ReadAll(file)
		file.Close()
	}

	if err := h.Service.UpdateTvSeries(id, &updateData); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse(w, http.StatusOK, map[string]string{
		"status":  "success",
		"message": "TV series updated",
	})
}

func (h *TvSeriesHandler) DeleteTvSeries(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id := uint(atoiSafe(idStr))

	if id == 0 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteTvSeries(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse(w, http.StatusOK, map[string]string{
		"status":  "success",
		"message": "TV series deleted",
	})
}

func atoiSafe(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
