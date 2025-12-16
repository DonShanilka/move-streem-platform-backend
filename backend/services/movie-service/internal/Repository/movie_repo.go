package repository

import (
	"github.com/DonShanilka/movie-service/internal/models"
	"gorm.io/gorm"
)

type MovieRepository struct {
	DB *gorm.DB
}

func (r *MovieRepository) Create(movie *models.Movie) error {
	return r.DB.Create(movie).Error
}

func (r *MovieRepository) Update(movie *models.Movie) error {
	return r.DB.Save(movie).Error
}

func (r *MovieRepository) FindAll() ([]models.Movie, error) {
	var movies []models.Movie
	err := r.DB.Find(&movies).Error
	return movies, err
}
