package models

type Movie struct {
	ID          int
	Title       string
	Description string
	ReleaseYear int
	Language    string
	Duration    int
	Rating      string
	AgeRating   string
	Country     string

	Thumbnail []byte
	Banner    []byte

	MovieURL string
	Trailer  []byte
}
