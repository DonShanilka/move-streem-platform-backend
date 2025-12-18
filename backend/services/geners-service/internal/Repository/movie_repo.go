package Repository

import (
	"github.com/DonShanilka/movie-service/internal/Models"
	"gorm.io/gorm"
)

type MovieRepository struct {
	DB *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{DB: db}
}

func (r *MovieRepository) Create(movie *Models.Movie) error {
	return r.DB.Create(movie).Error
}

func (r *MovieRepository) Update(id uint, movie *Models.Movie) error {
	return r.DB.Model(&Models.Movie{}).
		Where("id = ?", id).
		Updates(movie).Error
}

// Soft delete (industry standard)
func (r *MovieRepository) Delete(id uint) error {
	return r.DB.Delete(&Models.Movie{}, id).Error
}

func (r *MovieRepository) GetAll() ([]Models.Movie, error) {
	var movies []Models.Movie
	err := r.DB.Find(&movies).Error
	return movies, err
}
