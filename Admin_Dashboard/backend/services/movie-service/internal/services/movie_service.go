package services

import (
    "database/sql"
    "fmt"

    "github.com/DonShanilka/movie-service/internal/models"
    "github.com/DonShanilka/movie-service/internal/repository"
)

type MovieService struct {
    Repo *repository.MovieRepository
}

func NewMovieService(db *sql.DB) *MovieService {
    return &MovieService{
        Repo: repository.NewMovieRepository(db),
    }
}

func (s *MovieService) SaveMovie(movie models.Movie) error {
    return s.Repo.SaveMovie(movie)
}

func (s *MovieService) GetAllMovies() ([]models.Movie, error) {
    movies, err := s.Repo.GetAllMovies()
    if err != nil {
        return nil, err
    }

    for i := range movies {
        // Remove file
        movies[i].File = nil

        // Add video URL
        movies[i].VideoURL = fmt.Sprintf("http://localhost:8080/api/movies/stream?id=%d", movies[i].ID)
    }

    return movies, nil
}


func (s *MovieService) GetMovieFile(id int) ([]byte, error) {
    return s.Repo.GetMovieFile(id)
}
