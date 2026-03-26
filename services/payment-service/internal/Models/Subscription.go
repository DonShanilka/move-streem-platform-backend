package Models

import "time"

type Subscription struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	// 🔗 Foreign key (user-service, no relation struct needed)
	UserID uint `gorm:"not null;index"`

	// 🔗 Foreign key to plans table (local relation)
	PlanID uint `gorm:"not null;index"`
	Plan   Plan `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	// 📅 Dates
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`

	// 📌 Status
	Status string `gorm:"type:enum('active','expired','cancelled','paused');default:'active'"`

	// 💰 Payment
	Amount float64 `gorm:"not null"`

	// 🕒 Timestamps
	CreatedAt time.Time
	UpdatedAt time.Time
}
