package controllers

import (
		"github.com/DonShanilka/movie-service/internal/models"
		"github.com/DonShanilka/movie-service/internal/service"
)

type EpisodeController struct {
	Service *services.EpisodeService
}

func NewEpisodeController(service *services.EpisodeService) *EpisodeController {
	return &EpisodeController{Service: service}
}

func (c *EpisodeController) CreateEpisode(episode models.Episode) error {
	return c.Service.SaveEpisode(episode)
}

