package services

import (
    "database/sql"
    "github.com/DonShanilka/movie-service/internal/models"
)

type MovieService struct {
    DB *sql.DB
}

func NewMovieService(db *sql.DB) *MovieService {
    return &MovieService{DB: db}
}

func (s *MovieService) SaveMovie(movie models.Movie) error {
    query := `
        INSERT INTO movies (title, description, genre, release_year, duration, file)
        VALUES (?, ?, ?, ?, ?, ?)
    `
    _, err := s.DB.Exec(query,
        movie.Title,
        movie.Description,
        movie.Genre,
        movie.ReleaseYear,
        movie.Duration,
        movie.File,
    )
    return err
}


func (s *MovieService) GetAllMovies() ([]models.Movie, error) {
    rows, err := s.DB.Query("SELECT id, title, description, genre, release_year, duration, file FROM movies")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var movies []models.Movie

    for rows.Next() {
        var m models.Movie
        if err := rows.Scan(&m.ID, &m.Title, &m.Description, &m.Genre, &m.ReleaseYear, &m.Duration, &m.File); err != nil {
            return nil, err
        }
        movies = append(movies, m)
    }
    return movies, nil
}
