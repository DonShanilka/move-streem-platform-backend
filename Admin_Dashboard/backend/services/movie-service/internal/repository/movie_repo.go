package repository

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
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

    query := `CREATE TABLE IF NOT EXISTS movies (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255),
        description TEXT,
        genre VARCHAR(100),
        release_year YEAR,
        duration INT,
        file LONGBLOB
    );`

    _, err = db.Exec(query)
    if err != nil {
        return nil, err
    }

    fmt.Println("Database connected")
    return db, nil
}
