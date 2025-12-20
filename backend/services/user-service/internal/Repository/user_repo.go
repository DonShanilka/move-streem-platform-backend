package Repository

import (
	"github.com/DonShanilka/user-service/internal/Models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user *Models.User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepository) UpdateUser(id uint, user *Models.User) error {
	return repo.DB.Model(&Models.User{}).Where("id = ?", id).Updates(user).Error
}

func (repo *UserRepository) DeleteUser(id uint) error {
	return repo.DB.Delete(&Models.User{}, id).Error
}

func (repo *UserRepository) GetAllUser() ([]Models.User, error) {
	var users []Models.User
	err := repo.DB.Find(&users).Error
	return users, err
}
