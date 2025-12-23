package Models

type Series struct {
	ID          int
	Title       string
	Description string
	ReleaseYear int
	Language    string
	SeasonCount int
	Country     string
	AgeRating   string
	Rating      float64
	Genre       string
	Banner      []byte
	Trailer     []byte
}
