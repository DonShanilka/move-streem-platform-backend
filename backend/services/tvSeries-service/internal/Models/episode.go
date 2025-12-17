package Models

type Episode struct {
	ID            int
	SeriesID      int
	SeasonNumber  int
	EpisodeNumber int
	Title         string
	Description   string
	Duration      int
	ThumbnailURL  string
	Episode       string
	ReleaseDate   string
}
