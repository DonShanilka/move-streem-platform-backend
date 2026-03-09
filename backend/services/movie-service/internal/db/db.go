package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}

	dsn := fmt.Sprintf("root:Shanilka800@#@tcp(%s:%s)/movies_db?parseTime=true", dbHost, dbPort)
	log.Println("Database Connect Don ✅")
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
