package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/DonShanilka/movie-service/internal/models"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

	dsn := "root:Shanilka800@#@tcp(127.0.0.1:3306)/movies_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AUTOMIGRATE ALL MODELS
	err = db.AutoMigrate(
		&models.Movie{},
		// &models.Genre{},
		// &models.CastMember{},
		// &models.MovieGenre{},
		// &models.MovieCast{},
		// &models.Series{},
		// &models.SeriesGenre{},
		// &models.SeriesCast{},
		&models.Episode{},
		// &models.User{},
		// &models.Favorite{},
		// &models.WatchHistory{},
	)
	if err != nil {
		return nil, err
	}

	log.Println("✅ GORM: Database migrated successfully")
	DB = db
	return db, nil
}
