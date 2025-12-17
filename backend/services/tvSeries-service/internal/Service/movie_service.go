package services

import (
	"github.com/DonShanilka/movie-service/internal/Models"
	_ "github.com/DonShanilka/movie-service/internal/Models"
	"github.com/DonShanilka/movie-service/internal/Repository"
	_ "github.com/DonShanilka/movie-service/internal/Repository"
)

type MovieService struct {
	Repo *Repository.MovieRepository
}

func NewMovieService(repo *Repository.MovieRepository) *MovieService {
	return &MovieService{Repo: repo}
}

func (s *MovieService) CreateMovie(movie *Models.Movie) error {
	return s.Repo.Create(movie)
}

func (s *MovieService) UpdateMovie(id uint, movie *Models.Movie) error {
	return s.Repo.Update(id, movie)
}

func (s *MovieService) DeleteMovie(id uint) error {
	return s.Repo.Delete(id)
}

func (s *MovieService) GetAllMovies() ([]Models.Movie, error) {
	return s.Repo.GetAll()
}
