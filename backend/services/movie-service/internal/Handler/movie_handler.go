package Handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/DonShanilka/movie-service/internal/Models"
	services "github.com/DonShanilka/movie-service/internal/Service"
)

type MovieHandler struct {
	Service *services.MovieService
}

func NewMovieHandler(service *services.MovieService) *MovieHandler {
	return &MovieHandler{Service: service}
}

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(100 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("movie")
	if err != nil {
		http.Error(w, "Video File is Required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	movie := Models.Movie{
		Title:       r.FormValue("Title"),
		Description: r.FormValue("Description"),
		ReleaseYear: atoiSafe(r.FormValue("ReleaseYear")),
		Language:    r.FormValue("Language"),
		Duration:    atoiSafe(r.FormValue("Duration")),
		Rating:      r.FormValue("Rating"),
		AgeRating:   r.FormValue("AgeRating"),
		Country:     r.FormValue("Country"),
	}

	// Read files
	if file, _, _ := r.FormFile("Thumbnail"); file != nil {
		movie.Thumbnail, _ = io.ReadAll(file)
		file.Close()
	}
	if file, _, _ := r.FormFile("Banner"); file != nil {
		movie.Banner, _ = io.ReadAll(file)
		file.Close()
	}
	if file, _, _ := r.FormFile("Trailer"); file != nil {
		movie.Trailer, _ = io.ReadAll(file)
		file.Close()
	}

	err = h.Service.CreateMovie(&movie, file, header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Movie uploaded successfully",
	})
}

func atoiSafe(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ID
	idStr := r.FormValue("Id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	// ✅ Convert ReleaseYear
	releaseYearStr := r.FormValue("ReleaseYear")
	releaseYear, err := strconv.Atoi(releaseYearStr)
	if err != nil {
		http.Error(w, "Invalid ReleaseYear", http.StatusBadRequest)
		return
	}

	// ✅ Convert Duration
	durationStr := r.FormValue("Duration")
	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		http.Error(w, "Invalid Duration", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("movie")
	if err != nil {
		http.Error(w, "Video File is Required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	movie := Models.Movie{
		Id:          id,
		Title:       r.FormValue("Title"),
		Description: r.FormValue("Description"),
		ReleaseYear: releaseYear, // ✅ int
		Language:    r.FormValue("Language"),
		Duration:    duration, // ✅ int
		Rating:      r.FormValue("Rating"),
		AgeRating:   r.FormValue("AgeRating"),
		Country:     r.FormValue("Country"),
	}

	if err := h.Service.UpdateMovie(&movie, file, header.Filename); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Movie updated successfully",
	})
}

func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	if err := h.Service.DeleteMovie(uint(id)); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Movie deleted"})
}

func (h *MovieHandler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Service.GetAllMovies()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(movies)
}

func (h *MovieHandler) GetMovieById(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	if idstr == "" {
		http.Error(w, "Missing Movie ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid Movie ID", http.StatusBadRequest)
		return
	}

	movie, err := h.Service.GetMovieById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movie)
}
