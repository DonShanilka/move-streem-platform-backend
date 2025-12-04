package repository

import (
	"database/sql"
	"errors"

	"github.com/DonShanilka/movie-service/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type SeriesRepository struct {
	DB *sql.DB
}

func (r *MovieRepository) SaveSeries(series models.Se)