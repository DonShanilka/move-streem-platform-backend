package repository

import (
	"database/sql"
	// "errors"

	"github.com/DonShanilka/movie-service/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type SeriesRepository struct {
	DB *sql.DB
}

func (r *SeriesRepository) SaveSeries(series models.Series) error {
	quary := `INSERT INTO series (title, description, release_year, language, 					season_count, thumbnail_url, banner) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.DB.Exec(quary,
		series.Title,
		series.Description,
		series.ReleaseYear,
		series.Language,
		series.SeasonCount,
		series.ThumbnailURL,
		series.Banner,
	)
	return err

}
