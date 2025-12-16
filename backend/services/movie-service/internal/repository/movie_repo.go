package repository

import (
	"database/sql"
	"errors"

	"github.com/DonShanilka/movie-service/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type MovieRepository struct {
	DB *sql.DB
}

func (r *MovieRepository) SaveMovie(movie models.Movie) error {

	query := `INSERT INTO movies 
        (title, description, release_year, language, duration, rating,
         age_rating, country, thumbnail, banner, movie_url, trailer)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.DB.Exec(query,
		movie.Title,
		movie.Description,
		movie.ReleaseYear,
		movie.Language,
		movie.Duration,
		movie.Rating,
		movie.AgeRating,
		movie.Country,
		movie.Thumbnail,
		movie.Banner,
		movie.MovieURL,
		movie.Trailer,
	)
	return err
}

func (r *MovieRepository) UpdateMovie(id int, movie models.Movie) error {

	query := `
		UPDATE movies SET
			title = ?,
			description = ?,
			release_year = ?,
			language = ?,
			duration = ?,
			rating = ?,
			age_rating = ?,
			country = ?,
			thumbnail = ?,
			banner = ?,
			trailer = ?,
			movie_url = ?
		WHERE id = ?
	`

	_, err := r.DB.Exec(
		query,
		movie.Title,
		movie.Description,
		movie.ReleaseYear,
		movie.Language,
		movie.Duration,
		movie.Rating,
		movie.AgeRating,
		movie.Country,
		movie.Thumbnail,
		movie.Banner,
		movie.Trailer,
		movie.MovieURL,
		id,
	)

	return err
}

func (r *MovieRepository) GetAllMovies() ([]models.Movie, error) {
	query := `SELECT id, title, description, release_year, language, duration, rating,
         age_rating, country, thumbnail, banner, movie_url, trailer 
         FROM movies`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var m models.Movie
		err := rows.Scan(
			&m.ID,
			&m.Title,
			&m.Description,
			&m.ReleaseYear,
			&m.Language,
			&m.Duration,
			&m.Rating,
			&m.AgeRating,
			&m.Country,
			&m.Thumbnail,
			&m.Banner,
			&m.MovieURL,
			&m.Trailer,
		)
		if err != nil {
			return nil, err
		}

		movies = append(movies, m)
	}

	return movies, nil
}

func (r *MovieRepository) GetMovieFile(id int) ([]byte, error) {
	query := `SELECT file FROM movies WHERE id = ?`

	var fileData []byte
	err := r.DB.QueryRow(query, id).Scan(&fileData)

	if err == sql.ErrNoRows {
		return nil, errors.New("movie not found")
	}
	if err != nil {
		return nil, err
	}

	return fileData, nil
}
