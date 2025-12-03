package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/DonShanilka/movie-service/internal/models"
	_ "github.com/go-sql-driver/mysql"
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

	// --- Create movies table ---
	moviesTable := `
        CREATE TABLE movies (
        id INT PRIMARY KEY AUTO_INCREMENT,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        release_year INT,
        language VARCHAR(50),
        duration INT,
        rating DECIMAL(3,1),
        age_rating VARCHAR(10),
        country VARCHAR(100),
        thumbnail MEDIUMBLOB,
        banner MEDIUMBLOB,
        movie VARCHAR(255) NOT NULL,
        trailer LONGBLOB,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );`
	_, err = db.Exec(moviesTable)
	if err != nil {
		return nil, err
	}

	// --- Create genres table ---
	genresTable := `
    CREATE TABLE IF NOT EXISTS genres (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(50) NOT NULL UNIQUE
    );`
	_, err = db.Exec(genresTable)
	if err != nil {
		return nil, err
	}

    // --- Create cast_members table ---
	castTable := `CREATE TABLE IF NOT EXISTS cast_members (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    bio TEXT,
    avatar_url TEXT
    );`
	_, err = db.Exec(castTable)
	if err != nil {
		return nil, err
	}

    // movie_genres maping
    movieGenresTable := `CREATE TABLE IF NOT EXISTS movie_genres (
        movie_id INT,
        genre_id INT,
        PRIMARY KEY (movie_id, genre_id),
        FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE CASCADE,
        FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
    );`
    _, err = db.Exec(movieGenresTable)
    if err != nil {
        return nil, err
    }

    // movie_cast maping
    movieCastTable := `CREATE TABLE IF NOT EXISTS movie_cast (
        movie_id INT,
        cast_id INT,
        role VARCHAR(100),
        PRIMARY KEY (movie_id, cast_id),
        FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE CASCADE,
        FOREIGN KEY (cast_id) REFERENCES cast_members(id) ON DELETE CASCADE
    );`
    _, err = db.Exec(movieCastTable)
    if err != nil {
        return nil, err
    }

    // Tvshows table
    tvshowsTable := `CREATE TABLE IF NOT EXISTS series (
        id INT PRIMARY KEY AUTO_INCREMENT,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        release_year INT,
        language VARCHAR(50),
        season_count INT DEFAULT 1,
        thumbnail_url TEXT,
        banner_url TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );`
    _, err = db.Exec(tvshowsTable)
    if err != nil {
        return nil, err
    }

    // series_genres mapping table
    seriesGenresTable := `CREATE TABLE IF NOT EXISTS series_genres (
        series_id INT,
        genre_id INT,
        PRIMARY KEY (series_id, genre_id),
        FOREIGN KEY (series_id) REFERENCES series(id) ON DELETE CASCADE,
        FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
    );`
    _, err = db.Exec(seriesGenresTable)
    if err != nil {
       return nil, err
    }

    // Series–Cast Mapping
    seriesCastTable := `CREATE TABLE IF NOT EXISTS series_cast (
        series_id INT,
        cast_id INT,
        role VARCHAR(100),
        PRIMARY KEY (series_id, cast_id),
        FOREIGN KEY (series_id) REFERENCES series(id) ON DELETE CASCADE,
        FOREIGN KEY (cast_id) REFERENCES cast_members(id) ON DELETE CASCADE
    );`
    _, err = db.Exec(seriesCastTable)
    if err != nil {
       return nil, err
    }

    // Episodes Table
    episodesTable := `CREATE TABLE IF NOT EXISTS episodes (
    id INT PRIMARY KEY AUTO_INCREMENT,
    series_id INT NOT NULL,
    season_number INT NOT NULL,
    episode_number INT NOT NULL,
    title VARCHAR(255),
    description TEXT,
    duration INT,
    thumbnail_url TEXT,
    episode VARCHAR(255) NOT NULL,
    release_date DATE,
    FOREIGN KEY (series_id) REFERENCES series(id) ON DELETE CASCADE
);`
    _, err = db.Exec(episodesTable)
    if err != nil {
       return nil, err
    }

	fmt.Println("✔ Database initialized: movies + genres tables created")
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
