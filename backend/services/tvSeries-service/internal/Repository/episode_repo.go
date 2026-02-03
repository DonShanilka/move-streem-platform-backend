package Repository

import (
	"context"
	"errors"
	"io"
	"strings"

	"github.com/DonShanilka/tvSeries-service/internal/Models"
	"github.com/kurin/blazer/b2"
	"gorm.io/gorm"
)

type EpisodeRepository struct {
	DB     *gorm.DB
	B2     *b2.Client
	Bucket *b2.Bucket
}

// Extract file name from B2 URL
func extractB2FileName(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}

// Create new repo with B2
func NewEpisodeRepository(db *gorm.DB, keyID, appKey, bucketName string) (*EpisodeRepository, error) {
	ctx := context.Background()
	client, err := b2.NewClient(ctx, keyID, appKey)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(ctx, bucketName)
	if err != nil {
		return nil, err
	}
	return &EpisodeRepository{
		DB:     db,
		B2:     client,
		Bucket: bucket,
	}, nil
}

// Upload video to B2 and save metadata
func (r *EpisodeRepository) SaveEpisodeWithFile(ep *Models.Episode, file io.Reader, fileName string) error {
	ctx := context.Background()

	// Upload video to B2
	obj := r.Bucket.Object(fileName)
	writer := obj.NewWriter(ctx)
	if _, err := io.Copy(writer, file); err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}

	// Save B2 URL in episode field
	ep.Episode = obj.URL()

	// Save metadata to DB
	return r.DB.Create(ep).Error
}

// Update episode with new file
func (r *EpisodeRepository) UpdateEpisodeWithFile(ep *Models.Episode, file io.Reader, fileName string) error {
	ctx := context.Background()

	// Delete old video if exists
	if ep.Episode != "" {
		oldName := extractB2FileName(ep.Episode)
		obj := r.Bucket.Object(oldName)
		_ = obj.Delete(ctx) // ignore error if not found
	}

	// Upload new video
	obj := r.Bucket.Object(fileName)
	writer := obj.NewWriter(ctx)
	if _, err := io.Copy(writer, file); err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}

	// Update DB record
	ep.Episode = obj.URL()
	return r.DB.Save(ep).Error
}

// Delete episode and B2 file
func (r *EpisodeRepository) DeleteEpisode(id int) error {
	var ep Models.Episode

	if err := r.DB.First(&ep, id).Error; err != nil {
		return err
	}

	ctx := context.Background()
	if ep.Episode != "" {
		fileName := extractB2FileName(ep.Episode)
		obj := r.Bucket.Object(fileName)
		_ = obj.Delete(ctx)
	}

	return r.DB.Delete(&ep).Error
}

// Get all episodes
func (r *EpisodeRepository) GetAllEpisode() ([]Models.Episode, error) {
	var episodes []Models.Episode
	err := r.DB.Find(&episodes).Error
	return episodes, err
}

// Get episode by ID
func (r *EpisodeRepository) GetEpisodeByID(id int) (*Models.Episode, error) {
	var ep Models.Episode
	err := r.DB.First(&ep, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("episode not found")
		}
		return nil, err
	}
	return &ep, nil
}

// Get episodes by SeriesID (GORM version)
func (r *EpisodeRepository) GetEpisodesBySeriesID(seriesID int) ([]Models.Episode, error) {
	var episodes []Models.Episode

	err := r.DB.
		Where("series_id = ?", seriesID).
		Order("episode_number ASC").
		Find(&episodes).Error

	if err != nil {
		return nil, err
	}

	if len(episodes) == 0 {
		return nil, errors.New("no episodes found for this series")
	}

	return episodes, nil
}
