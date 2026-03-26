package Models

type MovieGenre struct {
	MovieID uint `gorm:"primaryKey"`
	GenreID uint `gorm:"primaryKey"`
}
