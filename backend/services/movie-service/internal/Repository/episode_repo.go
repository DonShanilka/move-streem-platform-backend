package repository

import (
	"database/sql"
	// "errors"

	"github.com/DonShanilka/movie-service/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type EpisodeRepository struct {
	DB *sql.DB
}

func (r *EpisodeRepository) SaveEpisode(episode models.Episode) error {
	quary := `INSERT INTO episodes 
		(series_id, season_number, episode_number, title, description, duration, thumbnail_url, episode, release_date)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.DB.Exec(quary, 
		episode.SeriesID, 
		episode.SeasonNumber, 
		episode.EpisodeNumber, 
		episode.Title, 
		episode.Description, 
		episode.Duration, 
		episode.ThumbnailURL, 
		episode.EpisodeURL, 
		episode.ReleaseDate,
	)
	return err
 // yupdate
}