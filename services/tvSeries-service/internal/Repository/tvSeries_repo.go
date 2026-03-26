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

func (r *TvSeriesRepository) GetAll() ([]Models.Series, error) {
	var series []Models.Series
	err := r.DB.
		Order("created_at DESC").
		Find(&series).Error
	return series, err
}

func (r *TvSeriesRepository) GetByID(id uint) (*Models.Series, error) {
	var series Models.Series
	err := r.DB.
		First(&series, id).
		Error

	if err != nil {
		return nil, err
	}

	return &series, nil
}

func (r *TvSeriesRepository) Update(id uint, data *Models.Series) error {
	return r.DB.
		Model(&Models.Series{}).
		Where("id = ?", id).
		Updates(data).
		Error
}

func (r *TvSeriesRepository) Delete(id uint) error {
	return r.DB.
		Delete(&Models.Series{}, id).
		Error
}
