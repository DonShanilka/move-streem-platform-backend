package models

import "time"

type Series struct {
    ID           int       `json:"id" db:"id"`
    Title        string    `json:"title" db:"title"`
    Description  string    `json:"description" db:"description"`
    ReleaseYear  *int      `json:"release_year" db:"release_year"`
    Language     *string   `json:"language" db:"language"`
    SeasonCount  *int      `json:"season_count" db:"season_count"`
    ThumbnailURL *string   `json:"thumbnail_url" db:"thumbnail_url"`
    Banner       []byte    `json:"banner" db:"banner"` // MEDIUMBLOB
    CreatedAt    time.Time `json:"created_at" db:"created_at"`
    UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
