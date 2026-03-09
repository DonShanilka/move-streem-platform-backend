package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/DonShanilka/user-service/internal/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetEnv helper to provide default values
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// 🔒 Central DB Config (no passing from main.go)
var (
	DB_USER = "root"
	DB_PASS = "Shanilka800@#"
	DB_HOST = GetEnv("DB_HOST", "localhost")
	DB_PORT = GetEnv("DB_PORT", "3306")
	DB_NAME = "movies_db"
)

// InitDB creates DB if not exists and connects GORM
func InitDB() (*gorm.DB, error) {

	// 1️⃣ Ensure database exists
	if err := createDatabaseIfNotExists(); err != nil {
		return nil, err
	}

	// 2️⃣ Connect GORM
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		DB_USER,
		DB_PASS,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Models.User{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connected successfully ✅")
	return db, nil
}

// 🔒 Internal helper (not used outside db package)
func createDatabaseIfNotExists() error {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/",
		DB_USER,
		DB_PASS,
		DB_HOST,
		DB_PORT,
	)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	_, err = sqlDB.Exec("CREATE DATABASE IF NOT EXISTS " + DB_NAME)
	if err != nil {
		return err
	}

	log.Println("Database ensured (exists or created) ✅")
	return nil
}
