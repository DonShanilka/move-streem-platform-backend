package Service

import (
	"github.com/DonShanilka/genres-service/internal/Models"
	"github.com/DonShanilka/genres-service/internal/Repository"
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

func (service *GenreService) UpdateGenre(id uint, genre *Models.Genre) error {
	return service.Repo.UpdateGenre(id, genre)
}

func (service *GenreService) DeleteGenre(id uint) error {
	return service.Repo.DeleteGenre(id)
}

func (service *GenreService) GetAllGenres() ([]Models.Genre, error) {
	return service.Repo.GetAllGenres()
}
