package Repository

import (
	"context"
	"errors"
	"time"

	"github.com/DonShanilka/movie-service/internal/Models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TvSeriesRepository struct {
	SeriesColl *mongo.Collection
}

func NewTvSeriesRepository(db *mongo.Database) *TvSeriesRepository {
	return &TvSeriesRepository{
		SeriesColl: db.Collection("series"),
	}
}

func (r *TvSeriesRepository) CreateSeries(series *Models.Series) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.SeriesColl.InsertOne(ctx, series)
	if err != nil {
		return primitive.NilObjectID, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("failed to convert inserted ID to ObjectID")
	}

	return id, nil
}
