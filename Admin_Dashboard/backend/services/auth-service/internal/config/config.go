package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	JWTSecret string
	MongoURI  string
	Database  string
}

func LoadConfig() (*Config, error) {
	godotenv.Load()
	
	config := &Config{
		Port:      os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		MongoURI:  os.Getenv("MONGO_URI"),
		Database:  os.Getenv("MONGO_DB"),
	}

	return config, nil
}
