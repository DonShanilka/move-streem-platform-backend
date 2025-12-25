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

	// Save B2 URL in correct field (matching MySQL)
	ep.Episode = obj.URL()

	// Save metadata to MySQL
	return r.DB.Create(ep).Error
}

func (r *EpisodeRepository) UpdateEpisodeWithFile(
	ep *Models.Episode,
	file io.Reader,
	fileName string,
) error {

	ctx := context.Background()

	// 1️⃣ Delete old video if exists
	if ep.Episode != "" {
		oldName := extractB2FileName(ep.Episode)
		obj := r.Bucket.Object(oldName)
		_ = obj.Delete(ctx) // ignore error if not found
	}

	// 2️⃣ Upload new video
	obj := r.Bucket.Object(fileName)
	writer := obj.NewWriter(ctx)
	if _, err := io.Copy(writer, file); err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}

	// 3️⃣ Update DB record
	ep.Episode = obj.URL()
	return r.DB.Save(ep).Error
}


func (r *EpisodeRepository) DeleteEpisode(id int) error {
	var ep Models.Episode

	// 1️⃣ Find episode
	if err := r.DB.First(&ep, id).Error; err != nil {
		return err
	}

	ctx := context.Background()

	// 2️⃣ Delete video from B2
	if ep.Episode != "" {
		fileName := extractB2FileName(ep.Episode)
		obj := r.Bucket.Object(fileName)
		_ = obj.Delete(ctx)
	}

	// 3️⃣ Delete DB record
	return r.DB.Delete(&ep).Error
}

func (r *EpisodeRepository) GetAllEpisode() ([]Models.Episode, error) {
	var episode []Models.Episode
	err := r.DB.Find(&episode).Error
	return episode, err
}

// Get episode by ID
func (r *EpisodeRepository) GetEpisodeByID(id int) (*Models.Episode, error) {
	var episode Models.Episode

	err := r.DB.First(&episode, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("episode not found")
		}
		return nil, err
	}

	return &episode, nil
}
