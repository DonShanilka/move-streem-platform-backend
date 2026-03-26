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

func (repo *AdminRepository) DeleteAdmin(id uint) error {
	return repo.DB.Delete(&Models.Admin{}, id).Error
}

func (repo *AdminRepository) GetAllAdmins() ([]Models.Admin, error) {
	var admins []Models.Admin
	err := repo.DB.Find(&admins).Error
	return admins, err
}
