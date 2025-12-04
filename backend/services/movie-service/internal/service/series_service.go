package services

import (
	"database/sql"

	"github.com/DonShanilka/movie-service/internal/models"
	"github.com/DonShanilka/movie-service/internal/repository"
)

type SeriesService struct {
	Repo *repository.SeriesRepository
}

func NewSeriesService(db *sql.DB) *SeriesService {
	return &SeriesService{
		Repo: &repository.SeriesRepository{DB: db},
	}
}

func (s *SeriesService) SaveSeries(series models.Series) error {
	return s.Repo.SaveSeries(series)
}