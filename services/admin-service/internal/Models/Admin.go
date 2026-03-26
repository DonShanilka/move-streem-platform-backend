package Models

type Admin struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"type:varchar(100);not null"`
	Email        string `gorm:"type:varchar(250);unique;not null"`
	Password     string `gorm:"type:text;not null"`
	ProfileImage []byte `gorm:"type:longblob"`
}
