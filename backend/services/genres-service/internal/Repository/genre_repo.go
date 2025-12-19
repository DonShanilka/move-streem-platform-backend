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

func (request *GenerRepostry) CreateGenre(genre *Models.Genre) error {
	return request.DB.Create(genre).Error
}

func (request *GenerRepostry) UpdateGenre(id uint, genre *Models.Genre) error {
	return request.DB.Model(&Models.Genre{}).Where("id =?", id).Updates(genre).Error
}

func (request *GenerRepostry) DeleteGenre(id uint) error {
	return request.DB.Delete(&Models.Genre{}, id).Error
}

func (request *GenerRepostry) GetAllGenres() ([]Models.Genre, error) {
	var genres []Models.Genre
	err := request.DB.Find(&genres).Error
	return genres, err
}
