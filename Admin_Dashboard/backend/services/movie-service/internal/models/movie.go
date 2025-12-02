package models

type Movie struct {
    ID          int     `json:"id"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Genre       string  `json:"genre"`
    ReleaseYear int     `json:"release_year"`
    Duration    int     `json:"duration"`
    File        []byte  `json:"-"`
    VideoURL    string  `json:"video_url"`
}


