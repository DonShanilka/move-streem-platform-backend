package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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
        movie_url VARCHAR(255) NOT NULL,
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

    // Users Table
    usersTable := `CREATE TABLE IF NOT EXISTS users (
        id INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(255),
        email VARCHAR(255) UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
    _, err = db.Exec(usersTable)
    if err != nil {
       return nil, err
    }

    // Favorites Table
    favoritesTable := `CREATE TABLE IF NOT EXISTS favorites (
        user_id INT,
        movie_id INT NULL,
        series_id INT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (user_id, movie_id, series_id),
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE CASCADE,
        FOREIGN KEY (series_id) REFERENCES series(id) ON DELETE CASCADE
    );`
    _, err = db.Exec(favoritesTable)
    if err != nil {
       return nil, err
    }

    // Watch History Table
    watchHistoryTable := `CREATE TABLE IF NOT EXISTS watch_history (
        id INT PRIMARY KEY AUTO_INCREMENT,
        user_id INT NOT NULL,
        movie_id INT NULL,
        episode_id INT NULL,
        progress INT CHECK (progress >= 0 AND progress <= 100),
        last_watch_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE CASCADE,
        FOREIGN KEY (episode_id) REFERENCES episodes(id) ON DELETE CASCADE
    );`
    _, err = db.Exec(watchHistoryTable)
    if err != nil {
       return nil, err
    }

	fmt.Println("✔ Database initialized: movies + genres tables created")
	return db, nil
}