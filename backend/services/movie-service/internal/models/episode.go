package models

type Episode struct {
    ID            int    
    SeriesID      int    
    SeasonNumber  int    
    EpisodeNumber int    
    Title         string 
    Description   string 
    Duration      int    
    ThumbnailURL  string 
    EpisodeURL    string      
    ReleaseDate   string   
}
