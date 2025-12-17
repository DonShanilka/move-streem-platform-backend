package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Episode represents a single episode
type Episode struct {
	EpisodeNumber int    `bson:"episode_number" json:"episode_number"`
	Title         string `bson:"title" json:"title"`
	Description   string `bson:"description" json:"description"`
	Duration      int    `bson:"duration" json:"duration"` // in minutes
	ThumbnailURL  string `bson:"thumbnail_url" json:"thumbnail_url"`
	EpisodeURL    string `bson:"episode_url" json:"episode_url"` // cloud video URL
	ReleaseDate   string `bson:"release_date" json:"release_date"`
}

// Season represents a season of a series
type Season struct {
	SeasonNumber int       `bson:"season_number" json:"season_number"`
	Episodes     []Episode `bson:"episodes" json:"episodes"`
}

// Series represents a TV series
type Series struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title        string             `bson:"title" json:"title"`
	Description  string             `bson:"description" json:"description"`
	ReleaseYear  int                `bson:"release_year" json:"release_year"`
	Language     string             `bson:"language" json:"language"`
	SeasonCount  int                `bson:"season_count" json:"season_count"`
	ThumbnailURL string             `bson:"thumbnail_url" json:"thumbnail_url"`
	BannerURL    string             `bson:"banner_url" json:"banner_url"`
	Seasons      []Season           `bson:"seasons" json:"seasons"`
}
