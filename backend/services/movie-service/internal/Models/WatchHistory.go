package models

import "time"

type WatchHistory struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"index;not null"`
	MovieID   *uint
	EpisodeID *uint
	Progress  int `gorm:"check:progress >= 0 AND progress <= 100"`
	LastWatchTime time.Time
}