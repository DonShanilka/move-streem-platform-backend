package models

import "time"

type Series struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"size:255;not null"`
	Description string
	ReleaseYear int
	Language    string `gorm:"size:50"`
	SeasonCount int    `gorm:"default:1"`

	ThumbnailURL string
	Banner       []byte `gorm:"type:MEDIUMBLOB"`

	Genres []Genre      `gorm:"many2many:series_genres"`
	Cast   []CastMember `gorm:"many2many:series_cast"`

	Episodes []Episode

	CreatedAt time.Time
	UpdatedAt time.Time
}