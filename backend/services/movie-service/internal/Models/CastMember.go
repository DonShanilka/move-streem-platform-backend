package models

type CastMember struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Bio       string
	AvatarURL string
}
