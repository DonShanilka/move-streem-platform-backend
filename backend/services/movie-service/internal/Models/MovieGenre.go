package models

type MovieGenre struct {
	MovieID uint `gorm:"primaryKey"`
	GenreID uint `gorm:"primaryKey"`
}