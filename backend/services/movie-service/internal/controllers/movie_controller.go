package controllers

import (
	"github.com/DonShanilka/movie-service/internal/models"
	"github.com/DonShanilka/movie-service/internal/service"
)

type MovieController struct {
	Service *services.MovieService
}

func NewMovieController(service *services.MovieService) *MovieController {
	return &MovieController{Service: service}
}

func (c *MovieController) CreateMovie(movie models.Movie) error {
	return c.Service.SaveMovie(movie)
}

func (c *MovieController) GetMovies() ([]models.Movie, error) {
	return c.Service.GetAllMovies()
}

func (c *MovieController) UpdateMovie(id int, movie models.Movie) error {
	return c.Service.UpdateMovie(id, movie)
}
