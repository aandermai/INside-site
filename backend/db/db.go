package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Инициализация БД
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "app.db")
	if err != nil {
		log.Fatal(err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT,
		last_name TEXT,
		email TEXT UNIQUE,
		password_hash TEXT
	);
	`
	if _, err := DB.Exec(query); err != nil {
		log.Fatal(err)
	}
}
