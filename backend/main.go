package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type RegisterRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

// Инициализация БД
func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "app.db")
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
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

// Обработчик регистрации
func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	// Обработка preflight запроса
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Некорректные данные"}`, http.StatusBadRequest)
		return
	}

	if req.Password != req.RepeatPassword {
		http.Error(w, `{"error":"Пароли не совпадают"}`, http.StatusBadRequest)
		return
	}

	if len(req.Password) < 8 {
		http.Error(w, `{"error":"Пароль слишком короткий"}`, http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, `{"error":"Ошибка сервера"}`, http.StatusInternalServerError)
		return
	}

	_, err = db.Exec(
		"INSERT INTO users (first_name, last_name, email, password_hash) VALUES (?, ?, ?, ?)",
		req.FirstName, req.LastName, req.Email, string(hash),
	)
	if err != nil {
		log.Println("Ошибка вставки в БД:", err)
		http.Error(w, `{"error":"Email уже используется или другая ошибка"}`, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func main() {
	initDB()

	http.HandleFunc("/register", registerHandler)

	log.Println("Сервер запущен на порту 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
