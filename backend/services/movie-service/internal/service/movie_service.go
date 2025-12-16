package services

import (
	"database/sql"

	"github.com/DonShanilka/movie-service/internal/models"
	"github.com/DonShanilka/movie-service/internal/repository"
)

type MovieService struct {
	Repo *repository.MovieRepository
}

func NewMovieService(db *sql.DB) *MovieService {
	return &MovieService{
		Repo: &repository.MovieRepository{DB: db},
	}
}

func (s *MovieService) SaveMovie(movie models.Movie) error {
	return s.Repo.SaveMovie(movie)
}

func (s *MovieService) UpdateMovie(id int, movie models.Movie) error {
	return s.Repo.UpdateMovie(id, movie)
}

func (s *MovieService) GetAllMovies() ([]models.Movie, error) {
	movies, err := s.Repo.GetAllMovies()
	if err != nil {
		return nil, err
	}

	for i := range movies {
		movies[i].MovieURL = "http://localhost:8080/movies/getAllMovies/" + movies[i].MovieURL
	}

	return movies, nil
}

// Stream main movie file (stored locally)
func (s *MovieService) GetMovieFile(id int) ([]byte, error) {
	return s.Repo.GetMovieFile(id)
}
