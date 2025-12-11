package models

import "time"

type Movie struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"size:255;not null"`
	Description string
	ReleaseYear int
	Language    string `gorm:"size:50"`
	Duration    int
	Rating      float32 `gorm:"type:decimal(3,1)"`
	AgeRating   string  `gorm:"size:10"`
	Country     string  `gorm:"size:100"`

	Thumbnail []byte `gorm:"type:MEDIUMBLOB"`
	Banner    []byte `gorm:"type:MEDIUMBLOB"`

	MovieURL string `gorm:"size:255;not null"`
	Trailer  []byte `gorm:"type:LONGBLOB"`

	Genres []Genre     `gorm:"many2many:movie_genres"`
	Cast   []CastMember `gorm:"many2many:movie_cast"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
