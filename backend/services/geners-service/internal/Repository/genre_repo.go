package Repository

import (
	"github.com/DonShanilka/geners-service/internal/Models"
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
