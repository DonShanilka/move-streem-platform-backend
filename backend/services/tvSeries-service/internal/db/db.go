package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

func InitMongoDB() (*mongo.Database, error) {
	// MongoDB Atlas connection URI
	uri := "mongodb+srv://N. Shanilka:Shanilka1@#@cluster0.mongodb.net/movies_db?retryWrites=true&w=majority"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("MongoDB connection failed ❌", err)
		return nil, err
	}

	// Ping to ensure connection is established
	if err := client.Ping(ctx, nil); err != nil {
		log.Println("MongoDB ping failed ❌", err)
		return nil, err
	}

	log.Println("MongoDB Atlas connected ✅")

	// Assign global client
	MongoClient = client
	MongoDB = client.Database("movies_db") // use the same database name

	return MongoDB, nil
}

// Helper function to get collection
func GetCollection(collectionName string) *mongo.Collection {
	return MongoDB.Collection(collectionName)
}
