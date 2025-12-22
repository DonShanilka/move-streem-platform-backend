package services

import (
	"io"

	"github.com/DonShanilka/movie-service/internal/Models"
	"github.com/DonShanilka/movie-service/internal/Repository"
)

type MovieService struct {
	Repo *Repository.MovieRepository
}

func NewMovieService(repo *Repository.MovieRepository) *MovieService {
	return &MovieService{Repo: repo}
}

func (s *MovieService) CreateMovie(movie *Models.Movie, file io.Reader, fileName string) error {
	return s.Repo.SaveMovieWithFile(movie, file, fileName)
}

func (s *MovieService) UpdateMovie(movie *Models.Movie, file io.Reader, fileName string) error {
	return s.Repo.UpdateMovieWithFile(movie, file, fileName)
}

func (s *MovieService) DeleteMovie(id uint) error {
	return s.Repo.DeleteMovie(id)
}

func (s *MovieService) GetAllMovies() ([]Models.Movie, error) {
	return s.Repo.GetAllMovie()
}

func (s *MovieService) GetMovieById(id uint) (*Models.Movie, error) {
	return s.Repo.GetMovieByID(id)
}
