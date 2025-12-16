package models

type MovieCast struct {
	MovieID uint   `gorm:"primaryKey"`
	CastID  uint   `gorm:"primaryKey"`
	Role    string `gorm:"size:100"`
}