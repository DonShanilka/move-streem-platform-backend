package services

import (
	"database/sql"

	"github.com/DonShanilka/movie-service/internal/Models"
	"github.com/DonShanilka/movie-service/internal/Repository"
)

type SeriesService struct {
	Repo *Repository.SeriesRepository
}

func NewSeriesService(db *sql.DB) *SeriesService {
	return &SeriesService{
		Repo: &Repository.SeriesRepository{DB: db},
	}
}

func (s *SeriesService) SaveSeries(series Models.Series) error {
	return s.Repo.SaveSeries(series)
}
