package models

type Genre struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:50;unique;not null"`
}
