package Models

type Plan struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"type:varchar(50);unique"`
	Price        float64
	DurationDays int
	Quality      string
	MaxDevices   int
}
