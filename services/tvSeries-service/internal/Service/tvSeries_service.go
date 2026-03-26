package Service

import (
	"errors"

	"github.com/DonShanilka/tvSeries-service/internal/Models"
	"github.com/DonShanilka/tvSeries-service/internal/Repository"
	"gorm.io/gorm"
)

type TvSerriesService struct {
	Repo *Repository.TvSeriesRepository
}

func NewTvSerriesService(repo *Repository.TvSeriesRepository) *TvSerriesService {
	return &TvSerriesService{Repo: repo}
}

func (s *TvSerriesService) CreateTvSeries(tvSeries *Models.Series) error {
	return s.Repo.Create(tvSeries)
}

func (s *TvSerriesService) GetAllTvSeries() ([]Models.Series, error) {
	return s.Repo.GetAll()
}

func (s *TvSerriesService) GetTvSeriesByID(id uint) (*Models.Series, error) {
	series, err := s.Repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tv series not found")
		}
		return nil, err
	}
	return series, nil
}

func (s *TvSerriesService) UpdateTvSeries(id uint, data *Models.Series) error {
	// Ensure series exists before update
	_, err := s.Repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("tv series not found")
		}
		return err
	}

	return s.Repo.Update(id, data)
}

func (s *TvSerriesService) DeleteTvSeries(id uint) error {
	// Ensure series exists before delete
	_, err := s.Repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("tv series not found")
		}
		return err
	}

	return s.Repo.Delete(id)
}
