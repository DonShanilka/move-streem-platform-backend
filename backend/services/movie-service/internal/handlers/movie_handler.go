package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/DonShanilka/movie-service/internal/models"
	"github.com/DonShanilka/movie-service/internal/service"
)

type MovieHandler struct {
	Service *services.MovieService
}

func NewMovieHandler(service *services.MovieService) *MovieHandler {
	return &MovieHandler{Service: service}
}

func (h *MovieHandler) UploadMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// max upload size: 100GB
	r.ParseMultipartForm(110 << 30)

	// ========== READ NORMAL FIELDS ==========
	title := r.FormValue("title")
	description := r.FormValue("description")
	releaseYear := atoiSafe(r.FormValue("release_year"))
	language := r.FormValue("language")
	duration := atoiSafe(r.FormValue("duration"))
	rating := r.FormValue("rating")
	ageRating := r.FormValue("age_rating")
	country := r.FormValue("country")

	// ========== READ THUMBNAIL (BLOB) ==========
	thumbFile, _, _ := r.FormFile("thumbnail")
	var thumbnail []byte
	if thumbFile != nil {
		thumbnail, _ = io.ReadAll(thumbFile)
		thumbFile.Close()
	}

	// ========== READ BANNER (BLOB) ==========
	bannerFile, _, _ := r.FormFile("banner")
	var banner []byte
	if bannerFile != nil {
		banner, _ = io.ReadAll(bannerFile)
		bannerFile.Close()
	}

	// ========== READ TRAILER (BLOB) ==========
	trailerFile, _, _ := r.FormFile("trailer")
	var trailer []byte
	if trailerFile != nil {
		trailer, _ = io.ReadAll(trailerFile)
		trailerFile.Close()
	}

	// ========== SAVE FULL MOVIE TO LOCAL DISK ==========
	movieFile, movieHeader, err := r.FormFile("movie")
	if err != nil {
		http.Error(w, "Movie file missing: "+err.Error(), 400)
		return
	}
	defer movieFile.Close()

	moviePath := "./movies/" + movieHeader.Filename

	os.MkdirAll("./movies", 0755)
	f, err := os.Create(moviePath)
	if err != nil {
		http.Error(w, "Error saving movie locally: "+err.Error(), 500)
		return
	}
	io.Copy(f, movieFile)
	f.Close()

	// ========== CREATE MODEL ==========
	movie := models.Movie{
		Title:       title,
		Description: description,
		ReleaseYear: releaseYear,
		Language:    language,
		Duration:    duration,
		Rating:      rating,
		AgeRating:   ageRating,
		Country:     country,
		Thumbnail:   thumbnail,
		Banner:      banner,
		Trailer:     trailer,
		MovieURL:    movieHeader.Filename, // only file name saved
	}

	// Save to DB
	if err := h.Service.SaveMovie(movie); err != nil {
		http.Error(w, "DB error: "+err.Error(), 500)
		return
	}

	jsonResponse(w, map[string]interface{}{
		"message":     "Movie uploaded successfully",
		"movie_local": moviePath,
	})
}

// Update movie
func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	r.ParseMultipartForm(110 << 30)

	movie := models.Movie{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		ReleaseYear: atoiSafe(r.FormValue("release_year")),
		Language:    r.FormValue("language"),
		Duration:    atoiSafe(r.FormValue("duration")),
		Rating:      r.FormValue("rating"),
		AgeRating:   r.FormValue("age_rating"),
		Country:     r.FormValue("country"),
	}

	// Thumbnail
	if f, _, _ := r.FormFile("thumbnail"); f != nil {
		movie.Thumbnail, _ = io.ReadAll(f)
		f.Close()
	}

	// Banner
	if f, _, _ := r.FormFile("banner"); f != nil {
		movie.Banner, _ = io.ReadAll(f)
		f.Close()
	}

	// Trailer
	if f, _, _ := r.FormFile("trailer"); f != nil {
		movie.Trailer, _ = io.ReadAll(f)
		f.Close()
	}

	if err := h.Service.UpdateMovie(id, movie); err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(w, map[string]string{
		"message": "Movie updated successfully",
	})
}

// List metadata only
func (h *MovieHandler) ListMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Service.GetAllMovies()
	if err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, movies)
}

// Stream movie by id
func (h *MovieHandler) StreamMovie(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(idStr)

	fileBytes, err := h.Service.GetMovieFile(id)
	if err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Length", strconv.Itoa(len(fileBytes)))
	w.Write(fileBytes)
}

// Helpers
func atoiSafe(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
