package Repository

import (
	"github.com/DonShanilka/genres-service/internal/Models"
	"gorm.io/gorm"
)

// Genre Database Manager
// hold databse connection
type GenerRepostry struct {
	DB *gorm.DB
}

// Creates a ready-to-use repository
// Injects the database dependency once
func NewGenerRepostry(db *gorm.DB) *GenerRepostry {
	return &GenerRepostry{DB: db}
}

func (repo *GenerRepostry) CreateGenre(genre *Models.Genre) error {
	return repo.DB.Create(genre).Error
}

func (repo *GenerRepostry) UpdateGenre(id uint, genre *Models.Genre) error {
	return repo.DB.Model(&Models.Genre{}).Where("id =?", id).Updates(genre).Error
}

func (repo *GenerRepostry) DeleteGenre(id uint) error {
	return repo.DB.Delete(&Models.Genre{}, id).Error
}

func (repo *GenerRepostry) GetAllGenres() ([]Models.Genre, error) {
	var genres []Models.Genre
	err := repo.DB.Find(&genres).Error
	return genres, err
}
