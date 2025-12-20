package Models

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"type:varchar(100);not null"`
	Email        string `gorm:"type:varchar(250);unique;not null"`
	Password     string `gorm:"type:text;not null"`
	ProfileImage []byte `gorm:"type:longblob"`
	IsActive     bool   `gorm:"default:true"`
	CreatedAt    time.Time
}
