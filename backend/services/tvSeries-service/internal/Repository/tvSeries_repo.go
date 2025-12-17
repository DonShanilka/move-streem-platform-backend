package Repository

import (
	"context"
	"errors"
	"time"

	"github.com/DonShanilka/tvSeries-service/internal/Models"
	"go.mongodb.org/mongo-driver/bson"
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

// Create Series
func (r *TvSeriesRepository) CreateSeries(series *Models.Series) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.SeriesColl.InsertOne(ctx, series)
	if err != nil {
		return primitive.NilObjectID, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("failed to convert inserted ID")
	}

	return id, nil
}

// Add Season to existing Series
func (r *TvSeriesRepository) AddSeasonToSeries(
	seriesID primitive.ObjectID,
	season Models.Season,
) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": seriesID}

	update := bson.M{
		"$push": bson.M{"seasons": season},
		"$inc":  bson.M{"season_count": 1},
	}

	_, err := r.SeriesColl.UpdateOne(ctx, filter, update)
	return err
}
