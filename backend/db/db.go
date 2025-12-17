package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB // Глобальное подключение к БД

func InitDB() {
	// Создаём папку data, если её нет
	os.MkdirAll("./data", 0755)

	// Открываем (или создаём) файл базы данных
	var err error
	DB, err = sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		log.Fatal("Не удалось открыть базу данных:", err)
	}

	// Проверяем, что подключение живое
	if err = DB.Ping(); err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}

	// Создаём таблицу users, если её нет
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT,
		last_name TEXT,
		email TEXT UNIQUE,
		password_hash TEXT
	);`

	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("Не удалось создать таблицу пользователей:", err)
	}

	log.Println("База данных готова к работе")
}
