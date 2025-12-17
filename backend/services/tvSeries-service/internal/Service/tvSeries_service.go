package Service

import (
	"errors"

	"github.com/DonShanilka/tvSeries-service/internal/Models"
	"github.com/DonShanilka/tvSeries-service/internal/Repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TvSeriesService struct {
	Repo *Repository.TvSeriesRepository
}

func NewTvSeriesService(
	repo *Repository.TvSeriesRepository,
) *TvSeriesService {
	return &TvSeriesService{Repo: repo}
}

// ---------------- CREATE SERIES ----------------
func (s *TvSeriesService) CreateSeries(
	series *Models.Series,
) (primitive.ObjectID, error) {

	if series.Title == "" {
		return primitive.NilObjectID, errors.New("title is required")
	}

	return s.Repo.CreateSeries(series)
}

// ---------------- ADD SEASON ----------------
func (s *TvSeriesService) AddSeason(
	seriesID primitive.ObjectID,
	season Models.Season,
) error {

	if season.SeasonNumber <= 0 {
		return errors.New("invalid season number")
	}

	return s.Repo.AddSeasonToSeries(seriesID, season)
}
