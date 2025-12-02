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

func (h *MovieHandler) UploadMovie(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Must parse multipart form
    err := r.ParseMultipartForm(50 << 30) // 50GB max
    if err != nil {
        http.Error(w, "Invalid form-data: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Read file
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "File error: "+err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Read file bytes INTO RAM
    fileBytes, err := io.ReadAll(file)
    if err != nil {
        http.Error(w, "File read error: "+err.Error(), http.StatusInternalServerError)
        return
    }

    movie := models.Movie{
        Title:       r.FormValue("title"),
        Description: r.FormValue("description"),
        Genre:       r.FormValue("genre"),
        ReleaseYear: atoiSafe(r.FormValue("release_year")),
        Duration:    atoiSafe(r.FormValue("duration")),
        File:        fileBytes,
    }

    if err := h.Service.SaveMovie(movie); err != nil {
        http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
        return
    }

    jsonResponse(w, map[string]string{
        "message": "Movie uploaded successfully",
        "file":    handler.Filename,
    })
}


func (h *MovieHandler) ListMovies(w http.ResponseWriter, r *http.Request) {
    movies, err := h.Service.GetAllMovies()
    if err != nil {
        http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    jsonResponse(w, movies)
}

func atoiSafe(s string) int {
    val, _ := strconv.Atoi(s)
    return val
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
