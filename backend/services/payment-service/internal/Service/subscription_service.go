package Service

import (
	"backend/payment-service/internal/Models"
	"backend/payment-service/internal/Repository"
)

type SubsService struct {
	Repo *Repository.SubsRepostry
}

func NewSubsService(repo *Repository.SubsRepostry) *SubsService {
	return &SubsService{Repo: repo}
}

func (service *SubsService) CreateSubs(subs *Models.Subscription) error {
	return service.Repo.CreateSubs(subs)
}

func (service *SubsService) UpdateSubs(id uint, subs *Models.Subscription) error {
	return service.Repo.UpdateSubs(id, subs)
}

func (service *SubsService) DeleteSubs(id uint) error {
	return service.Repo.DeleteSubs(id)
}

func (service *SubsService) GetAllSubs() ([]Models.Subscription, error) {
	return service.Repo.GetAllSubs()
}
