package Service

import (
	"github.com/DonShanilka/user-service/internal/Models"
	"github.com/DonShanilka/user-service/internal/Repository"
)

type UserService struct {
	Repo *Repository.UserRepository
}

func NewUserService(repo *Repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (service *UserService) CreateUser(admin *Models.User) error {
	return service.Repo.CreateUser(admin)
}

func (service *UserService) UpdateUser(id uint, admin *Models.User) error {
	return service.Repo.UpdateUser(id, admin)
}

func (service *UserService) DeleteUser(id uint) error {
	return service.Repo.DeleteUser(id)
}

func (service *UserService) GetAllUsers() ([]Models.User, error) {
	return service.Repo.GetAllUser()
}
