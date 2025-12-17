package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "root:Shanilka800@#@tcp(localhost:3306)/movies_db?parseTime=true"
	log.Println("Database Connect Don ✅")
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
