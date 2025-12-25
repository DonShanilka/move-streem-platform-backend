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
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		ReleaseYear: atoiSafe(r.FormValue("releaseYear")),
		Language:    r.FormValue("language"),
		Duration:    atoiSafe(r.FormValue("duration")),
		Rating:      r.FormValue("rating"),
		AgeRating:   r.FormValue("ageRating"),
		Country:     r.FormValue("country"),
		Genre:       r.FormValue("genre"),
	}

	// Read files
	if file, _, _ := r.FormFile("thumbnail"); file != nil {
		movie.Thumbnail, _ = io.ReadAll(file)
		file.Close()
	}
	if file, _, _ := r.FormFile("banner"); file != nil {
		movie.Banner, _ = io.ReadAll(file)
		file.Close()
	}
	if file, _, _ := r.FormFile("trailer"); file != nil {
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
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
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
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		ReleaseYear: atoiSafe(r.FormValue("releaseYear")),
		Language:    r.FormValue("language"),
		Duration:    atoiSafe(r.FormValue("duration")),
		Rating:      r.FormValue("rating"),
		AgeRating:   r.FormValue("ageRating"),
		Country:     r.FormValue("country"),
		Genre:       r.FormValue("genre"),
	}

	if file, _, _ := r.FormFile("thumbnail"); file != nil {
		movie.Thumbnail, _ = io.ReadAll(file)
		file.Close()
	}
	if file, _, _ := r.FormFile("banner"); file != nil {
		movie.Banner, _ = io.ReadAll(file)
		file.Close()
	}
	if file, _, _ := r.FormFile("trailer"); file != nil {
		movie.Trailer, _ = io.ReadAll(file)
		file.Close()
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

	// 🔥 IF "movie" QUERY EXISTS → STREAM VIDEO
	videoURL := r.URL.Query().Get("movie")
	if videoURL != "" {

		req, err := http.NewRequest("GET", videoURL, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Forward Range header (important for streaming)
		if rangeHeader := r.Header.Get("Range"); rangeHeader != "" {
			req.Header.Set("Range", rangeHeader)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Failed to fetch video", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		// Copy headers
		for k, v := range resp.Header {
			w.Header()[k] = v
		}

		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
		return
	}

	// 🔵 OTHERWISE → RETURN ALL MOVIES
	movies, err := h.Service.GetAllMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
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


