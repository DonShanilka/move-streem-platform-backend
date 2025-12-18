package services

import (
	"github.com/DonShanilka/geners-service/internal/Models"
	"github.com/DonShanilka/geners-service/internal/Repository"
)

type GenreService struct {
	Repo *Repository.GenerRepostry
}

func NewGenreService(repo *Repository.GenerRepostry) *GenreService {
	return &GenreService{Repo: repo}
}

func (service *GenreService) CreateGenre(genre *Models.Genre) error {
	return service.Repo.CreateGenre(genre)
}
