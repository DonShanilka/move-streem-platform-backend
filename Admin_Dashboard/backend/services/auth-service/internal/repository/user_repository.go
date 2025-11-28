package repository

import (
	"context"
	"github.com/DonShanilka/auth-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(col *mongo.Collection) *UserRepository {
	return &UserRepository{Collection: col}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.Collection.InsertOne(context.TODO(), user)
	return err
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.Collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return &user, err
}
