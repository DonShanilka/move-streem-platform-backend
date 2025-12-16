package controllers

import (
	"github.com/DonShanilka/movie-service/internal/models"
	"github.com/DonShanilka/movie-service/internal/service"
)

type SeriesController struct {
	Service *services.SeriesService
}

func NewSeriesController(service *services.SeriesService) *SeriesController {
	return &SeriesController{Service: service}
}

func (c *SeriesController) CreateSeries(service models.Series) error {
	return c.Service.SaveSeries(service)
}