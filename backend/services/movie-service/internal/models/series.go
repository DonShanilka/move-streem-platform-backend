package models


type Series struct {
    ID           int       
    Title        string    
    Description  string    
    ReleaseYear  int      
    Language     string  
    SeasonCount  int     
    ThumbnailURL string  
    Banner       []byte    
    
}