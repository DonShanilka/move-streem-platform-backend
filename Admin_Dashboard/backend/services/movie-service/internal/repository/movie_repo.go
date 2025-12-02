package repository

import (
    "database/sql"
    "errors"

    _ "github.com/go-sql-driver/mysql"
    "github.com/DonShanilka/movie-service/internal/models"
)

type MovieRepository struct {
    DB *sql.DB
}

func InitDB() (*sql.DB, error) {
    dsn := "root:Shanilka800@#@tcp(127.0.0.1:3306)/movies_db"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    // CREATE TABLE (run one time)
    createTable := `CREATE TABLE IF NOT EXISTS movies (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        genre VARCHAR(100),
        release_year INT,
        duration INT,
        file LONGBLOB
    )`
    _, err = db.Exec(createTable)
    if err != nil {
        return nil, err
    }

    return db, nil
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
    return &MovieRepository{DB: db}
}

func (r *MovieRepository) SaveMovie(movie models.Movie) error {

    query := `INSERT INTO movies 
        (title, description, genre, release_year, duration, file)
        VALUES (?, ?, ?, ?, ?, ?)`

    _, err := r.DB.Exec(query,
        movie.Title,
        movie.Description,
        movie.Genre,
        movie.ReleaseYear,
        movie.Duration,
        movie.File,
    )
    return err
}

func (r *MovieRepository) GetAllMovies() ([]models.Movie, error) {
    query := `SELECT id, title, description, genre, release_year, duration FROM movies`
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var movies []models.Movie

    for rows.Next() {
        var m models.Movie
        err := rows.Scan(
            &m.ID,
            &m.Title,
            &m.Description,
            &m.Genre,
            &m.ReleaseYear,
            &m.Duration,
        )
        if err != nil {
            return nil, err
        }

        movies = append(movies, m)
    }
    return movies, nil
}

func (r *MovieRepository) GetMovieFile(id int) ([]byte, error) {
    query := `SELECT file FROM movies WHERE id = ?`

    var fileData []byte
    err := r.DB.QueryRow(query, id).Scan(&fileData)

    if err == sql.ErrNoRows {
        return nil, errors.New("movie not found")
    }
    if err != nil {
        return nil, err
    }

    return fileData, nil
}
