package Repository

import (
	"github.com/DonShanilka/tvSeries-service/internal/Models"
	"gorm.io/gorm"
)

type TvSeriesRepository struct {
	DB *gorm.DB
}

func NewTvSeriesRepository(db *gorm.DB) *TvSeriesRepository {
	return &TvSeriesRepository{DB: db}
}

func (r *TvSeriesRepository) Create(tvSeries *Models.Series) error {
	return r.DB.Create(tvSeries).Error
}
