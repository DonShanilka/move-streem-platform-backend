package services

import (
	"database/sql"

	"github.com/DonShanilka/movie-service/internal/Models"
	"github.com/DonShanilka/movie-service/internal/Repository"
)

type EpisodeService struct {
	Repo *Repository.EpisodeRepository
}

func NewEpisodeService(db *sql.DB) *EpisodeService {
	return &EpisodeService{
		Repo: &Repository.EpisodeRepository{DB: db},
	}
}

func (s *EpisodeService) SaveEpisode(episode Models.Episode) error {
	return s.Repo.SaveEpisode(episode)
}
