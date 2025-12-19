package Repository

import (
	"github.com/DonShanilka/admin-service/internal/Models"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{DB: db}
}

func (repo *AdminRepository) CreateAdmin(admin *Models.Admin) error {
	return repo.DB.Create(admin).Error
}

func (repo *AdminRepository) UpdateAdmin(id uint, admin *Models.Admin) error {
	return repo.DB.Model(&Models.Admin{}).Where("id = ?", id).Updates(admin).Error
}
