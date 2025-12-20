package Repository

import (
	"backend/payment-service/internal/Models"
	"gorm.io/gorm"
)

type PlanRepostry struct {
	DB *gorm.DB
}

func NewPlanRepostry(db *gorm.DB) *PlanRepostry {
	return &PlanRepostry{DB: db}
}

func (repo *PlanRepostry) CreatePlan(plan *Models.Plan) error {
	return repo.DB.Create(plan).Error
}

func (repo *PlanRepostry) UpdatePlan(id uint, plan *Models.Plan) error {
	return repo.DB.Model(&Models.Plan{}).Where("id =?", id).Updates(plan).Error
}

func (repo *PlanRepostry) DeletePlan(id uint) error {
	return repo.DB.Delete(&Models.Plan{}, id).Error
}

func (repo *PlanRepostry) GetAllPlan() ([]Models.Plan, error) {
	var plane []Models.Plan
	err := repo.DB.Find(&plane).Error
	return plane, err
}
