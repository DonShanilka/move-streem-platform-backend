package models

import "time"

type Favorite struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint `gorm:"index;not null"`
	MovieID  *uint
	SeriesID *uint
	CreatedAt time.Time
}