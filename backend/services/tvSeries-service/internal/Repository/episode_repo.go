package Repository

import (
	"context"
	"io"

	"github.com/DonShanilka/tvSeries-service/internal/Models"
	"github.com/kurin/blazer/b2"
	"gorm.io/gorm"
)

type EpisodeRepository struct {
	DB     *gorm.DB
	B2     *b2.Client
	Bucket *b2.Bucket
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

	// Save B2 URL in correct field (matching MySQL)
	ep.Episode = obj.URL()

	// Save metadata to MySQL
	return r.DB.Create(ep).Error
}
