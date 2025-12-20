package Models

import (
	"time"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

type Subscription struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"` // FK column
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PlanID    uint
	StartDate time.Time
	EndDate   time.Time
	Status    string `gorm:"type:varchar(20)"`
	Amount    float64
	CreatedAt time.Time
}
