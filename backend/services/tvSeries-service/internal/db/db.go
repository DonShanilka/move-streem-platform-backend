package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

func InitMongoDB() (*mongo.Database, error) {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	uri := os.Getenv("MONGO_URI") // load from env
	if uri == "" {
		log.Fatal("MONGO_URI is empty! Check your .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("MongoDB connection failed ❌", err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Println("MongoDB ping failed ❌", err)
		return nil, err
	}

	log.Println("MongoDB Atlas connected ✅")

	MongoClient = client
	MongoDB = client.Database("movies_db")
	return MongoDB, nil
}
