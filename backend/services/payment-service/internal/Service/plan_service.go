package Service

import (
	"backend/payment-service/internal/Models"
	"backend/payment-service/internal/Repository"
)

type PlanService struct {
	Repo *Repository.PlanRepostry
}

func NewPlanService(repo *Repository.PlanRepostry) *PlanService {
	return &PlanService{Repo: repo}
}

func (service *PlanService) CreateGenre(plan *Models.Plan) error {
	return service.Repo.CreatePlan(plan)
}

func (service *PlanService) UpdateGenre(id uint, plan *Models.Plan) error {
	return service.Repo.UpdatePlan(id, plan)
}

func (service *PlanService) DeleteGenre(id uint) error {
	return service.Repo.DeletePlan(id)
}

func (service *PlanService) GetAllGenres() ([]Models.Plan, error) {
	return service.Repo.GetAllPlan()
}
