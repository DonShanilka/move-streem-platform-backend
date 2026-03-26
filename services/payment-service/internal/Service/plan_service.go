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

func (service *PlanService) CreatePlan(plan *Models.Plan) error {
	return service.Repo.CreatePlan(plan)
}

func (service *PlanService) UpdatePlan(id uint, plan *Models.Plan) error {
	return service.Repo.UpdatePlan(id, plan)
}

func (service *PlanService) DeletePlan(id uint) error {
	return service.Repo.DeletePlan(id)
}

func (service *PlanService) GetAllPlan() ([]Models.Plan, error) {
	return service.Repo.GetAllPlan()
}
