package Repository

import (
	"backend/payment-service/internal/Models"

	"gorm.io/gorm"
)

type SubsRepostry struct {
	DB *gorm.DB
}

func NewSubsRepostry(db *gorm.DB) *SubsRepostry {
	return &SubsRepostry{DB: db}
}

func (repo *SubsRepostry) CreateSubs(subs *Models.Subscription) error {
	return repo.DB.Create(subs).Error
}

func (repo *SubsRepostry) UpdateSubs(id uint, subs *Models.Subscription) error {
	return repo.DB.Model(&Models.Subscription{}).Where("id =?", id).Updates(subs).Error
}

func (repo *SubsRepostry) DeleteSubs(id uint) error {
	return repo.DB.Delete(&Models.Subscription{}, id).Error
}

func (repo *SubsRepostry) GetAllSubs() ([]Models.Subscription, error) {
	var subs []Models.Subscription
	err := repo.DB.Find(&subs).Error
	return subs, err
}
