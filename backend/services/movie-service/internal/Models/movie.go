package Models

type Movie struct {
	Id          int
	Title       string
	Description string
	ReleaseYear int
	Language    string
	Duration    int
	Rating      string
	AgeRating   string
	Country     string
	Thumbnail   []byte `gorm:"type:longblob"`
	Banner      []byte `gorm:"type:longblob"`
	Trailer     []byte `gorm:"type:longblob"`
	MovieURL    string
}
