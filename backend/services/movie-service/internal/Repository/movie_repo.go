package Repository

import (
	"context"
	"errors"
	"io"
	"strings"

	"github.com/DonShanilka/movie-service/internal/Models"
	"github.com/kurin/blazer/b2"
	"gorm.io/gorm"
)

type MovieRepository struct {
	DB     *gorm.DB
	B2     *b2.Client
	Bucket *b2.Bucket
}

func extractB2FileName(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}

func NewMovieRepository(db *gorm.DB, keyID, appKey, bucketName string) (*MovieRepository, error) {
	ctx := context.Background()
	client, err := b2.NewClient(ctx, keyID, appKey)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(ctx, bucketName)
	if err != nil {
		return nil, err
	}
	return &MovieRepository{
		DB:     db,
		B2:     client,
		Bucket: bucket,
	}, nil
}

func (r *MovieRepository) SaveMovieWithFile(movie *Models.Movie, file io.Reader, fileName string) error {
	ctx := context.Background()

	// Upload video to B2 and metadata
	obj := r.Bucket.Object(fileName)
	writer := obj.NewWriter(ctx)
	if _, err := io.Copy(writer, file); err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}

	movie.MovieURL = obj.URL()
	return r.DB.Save(movie).Error
}

func (r *MovieRepository) UpdateMovieWithFile(movie *Models.Movie, file io.Reader, fileName string) error {

	ctx := context.Background()

	// Delete old video if exists
	if movie.MovieURL != "" {
		oldName := extractB2FileName(movie.MovieURL)
		obj := r.Bucket.Object(oldName)
		_ = obj.Delete(ctx)
	}

	obj := r.Bucket.Object(fileName)
	writer := obj.NewWriter(ctx)
	if _, err := io.Copy(writer, file); err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}

	movie.MovieURL = obj.URL()
	return r.DB.Save(movie).Error
}

// Soft delete (industry standard)
func (r *MovieRepository) DeleteMovie(id uint) error {
	var movie Models.Movie

	if err := r.DB.First(&movie, id).Error; err != nil {
		return err
	}

	ctx := context.Background()

	if movie.MovieURL != "" {
		fileName := movie.MovieURL
		obj := r.Bucket.Object(fileName)
		_ = obj.Delete(ctx)
	}
	return r.DB.Delete(movie).Error
}

func (r *MovieRepository) GetAllMovie() ([]Models.Movie, error) {
	var movies []Models.Movie
	err := r.DB.Find(&movies).Error
	return movies, err
}

func (r *MovieRepository) GetMovieByID(id int) (*Models.Movie, error) {
	var movie Models.Movie

	err := r.DB.First(&movie, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("movie not found")
		}
		return nil, err
	}
	return &movie, nil
}
