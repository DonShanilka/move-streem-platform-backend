package Repository

import (
	"github.com/DonShanilka/genres-service/internal/Models"
	"gorm.io/gorm"
)

type GenerRepostry struct {
	DB *gorm.DB
}

func NewGenerRepostry(db *gorm.DB) *GenerRepostry {
	return &GenerRepostry{DB: db}
}

func (r *GenerRepostry) CreateGenre(genre *Models.Genre) error {
	return r.DB.Create(genre).Error
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
