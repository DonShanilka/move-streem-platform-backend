package Service

import (
	"github.com/DonShanilka/admin-service/internal/Models"
	"github.com/DonShanilka/admin-service/internal/Repository"
)

type AdminService struct {
	Repo *Repository.AdminRepository
}

func NewAdminService(repo *Repository.AdminRepository) *AdminService {
	return &AdminService{Repo: repo}
}

func (service *AdminService) CreateAdmin(admin *Models.Admin) error {
	return service.Repo.CreateAdmin(admin)
}

func (service *AdminService) UpdateAdmin(id uint, admin *Models.Admin) error {
	return service.Repo.UpdateAdmin(id, admin)
}
