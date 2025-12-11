package models

import "time"

type Episode struct {
	ID            uint      `gorm:"primaryKey"`
	SeriesID      uint      `gorm:"index"`
	SeasonNumber  int       `gorm:"not null"`
	EpisodeNumber int       `gorm:"not null"`

	Title       string `gorm:"size:255"`
	Description string
	Duration    int
	ThumbnailURL string

	EpisodeURL  string `gorm:"size:255;not null"`

	ReleaseDate *time.Time
}
