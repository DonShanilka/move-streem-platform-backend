package Service

import (
	"github.com/DonShanilka/user-service/internal/Models"
	"github.com/DonShanilka/user-service/internal/Repository"
)

type UserService struct {
	Repo *Repository.UserRepository
}

func NewAdminService(repo *Repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (service *UserService) CreateAdmin(admin *Models.User) error {
	return service.Repo.CreateUser(admin)
}

func (service *UserService) UpdateAdmin(id uint, admin *Models.User) error {
	return service.Repo.UpdateUser(id, admin)
}

func (service *UserService) DeleteAdmin(id uint) error {
	return service.Repo.DeleteUser(id)
}

func (service *UserService) GetAllAdmins() ([]Models.User, error) {
	return service.Repo.GetAllUser()
}
