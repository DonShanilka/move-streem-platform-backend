package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dns := "root:Shanilka800@#@tcp(localhost:3306)/movies_db?parseTime=true"
	log.Printf("Database Connect Don ✅")
	return gorm.Open(mysql.Open(dns), &gorm.Config{})
}
