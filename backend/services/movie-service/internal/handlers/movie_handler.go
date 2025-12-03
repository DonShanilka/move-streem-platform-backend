package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
    "os"

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

    // Allow up to 2GB upload form (enough for local file)
    err := r.ParseMultipartForm(2 << 30)
    if err != nil {
        http.Error(w, "Error parsing multipart form: "+err.Error(), http.StatusBadRequest)
        return
    }

    // -------------------------
    // 1️⃣ Read Movie File (store locally)
    // -------------------------
    movieFile, movieHeader, err := r.FormFile("movie")
    if err != nil {
        http.Error(w, "Movie file required", http.StatusBadRequest)
        return
    }
    defer movieFile.Close()

    movieBytes, err := io.ReadAll(movieFile)
    if err != nil {
        http.Error(w, "Error reading movie file", http.StatusInternalServerError)
        return
    }

    // Save movie to local folder
    moviePath := "movies/" + movieHeader.Filename
    err = os.WriteFile(moviePath, movieBytes, 0644)
    if err != nil {
        http.Error(w, "Error saving movie locally", http.StatusInternalServerError)
        return
    }

    // -------------------------
    // 2️⃣ Trailer (Save to MySQL as LONGBLOB)
    // -------------------------
    trailerFile, _, err := r.FormFile("trailer")
    var trailerBytes []byte
    if err == nil { // trailer is optional
        defer trailerFile.Close()
        trailerBytes, _ = io.ReadAll(trailerFile)
    }

    // -------------------------
    // 3️⃣ Thumbnail (MEDIUMBLOB)
    // -------------------------
    thumbnailFile, _, err := r.FormFile("thumbnail")
    var thumbnailBytes []byte
    if err == nil {
        defer thumbnailFile.Close()
        thumbnailBytes, _ = io.ReadAll(thumbnailFile)
    }

    // -------------------------
    // 4️⃣ Banner (MEDIUMBLOB)
    // -------------------------
    bannerFile, _, err := r.FormFile("banner")
    var bannerBytes []byte
    if err == nil {
        defer bannerFile.Close()
        bannerBytes, _ = io.ReadAll(bannerFile)
    }

    movie := models.Movie{
        Title:       r.FormValue("title"),
        Description: r.FormValue("description"),
        ReleaseYear: atoiSafe(r.FormValue("release_year")),
        Duration:    atoiSafe(r.FormValue("duration")),
        Language:    r.FormValue("language"),
        Country:     r.FormValue("country"),
        Rating:      float64(atoiSafe(r.FormValue("rating"))),
        AgeRating:   r.FormValue("age_rating"),

        Thumbnail: thumbnailBytes,
        Banner:    bannerBytes,

        MovieURL: moviePath,
        Trailer:  trailerBytes,
    }

    if err := h.Service.SaveMovie(movie); err != nil {
        http.Error(w, "DB save error: "+err.Error(), http.StatusInternalServerError)
        return
    }

    jsonResponse(w, map[string]string{
        "message": "Movie uploaded successfully",
        "movie_path": moviePath,
        "trailer_size": strconv.Itoa(len(trailerBytes)),
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
