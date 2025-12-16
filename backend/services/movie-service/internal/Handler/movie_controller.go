package controllers

import (
	"net/http"
	"strconv"

	"github.com/DonShanilka/movie-service/internal/models"
	"github.com/DonShanilka/movie-service/internal/services"
)

type MovieController struct {
	Service *services.MovieService
}

func NewMovieController(service *services.MovieService) *MovieController {
	return &MovieController{Service: service}
}

// GET /api/movies
func (c *MovieController) GetAll(w http.ResponseWriter, r *http.Request) {
	movies, err := c.Service.GetAllMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, movies)
}

// POST /api/movies
func (c *MovieController) Create(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	if err := decodeJSON(r, &movie); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := c.Service.CreateMovie(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(w, map[string]string{"message": "Movie created"})
}

// PUT /api/movies/{id}
func (c *MovieController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var movie models.Movie
	if err := decodeJSON(r, &movie); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := c.Service.UpdateMovie(id, movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(w, map[string]string{"message": "Movie updated"})
}
