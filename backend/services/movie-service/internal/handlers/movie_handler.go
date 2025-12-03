package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/DonShanilka/movie-service/internal/models"
	"github.com/DonShanilka/movie-service/internal/services"
)

type MovieHandler struct {
    Service *services.MovieService
}

func NewMovieHandler(service *services.MovieService) *MovieHandler {
    return &MovieHandler{Service: service}
}

// Upload movie with file stored in DB
func (h *MovieHandler) UploadMovie(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse multipart form
    err := r.ParseMultipartForm(50 << 30) // 50GB max
    if err != nil {
        http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Read file
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error reading file: "+err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Read file content
    fileBytes, err := io.ReadAll(file)
    if err != nil {
        http.Error(w, "Error reading file content: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // create movie model
    movie := models.Movie {
        Title:       r.FormValue("title"),
        Description: r.FormValue("description"),
        Genre:       r.FormValue("genre"),
        ReleaseYear: atoiSafe(r.FormValue("release_year")),
        Duration:    atoiSafe(r.FormValue("duration")),
        File:        fileBytes,
    }

    // Save movie
    if err = h.Service.SaveMovie(movie); err != nil {
        http.Error(w, "DB error: ❌"+err.Error(), http.StatusInternalServerError)
        return
    }

    jsonResponse(w, map[string]string{
        "message":  "Movie '" + handler.Filename + "' uploaded successfully ✅",
        "file_size": strconv.Itoa(len(fileBytes)),
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
