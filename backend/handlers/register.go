package handlers

import (
	"INside-site/backend/db"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// Структура для принятия данных из JSON-запроса
type RegisterRequest struct {
	FirstName      string `json:"first_name"`      // Имя
	LastName       string `json:"last_name"`       // Фамилия
	Email          string `json:"email"`           // Почта
	Password       string `json:"password"`        // Пароль
	RepeatPassword string `json:"repeat_password"` // Повтора пароля
}

// Обработчик регистрации
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")              // Разрешение запросов с любого домена
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")  // Разрешение заголовка Content-Type
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS") // Разрешение методов POST and OPTIONS

	// Обработка preflight запроса (OPTIONS-запрос)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Проверка POST-запроса
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest // Переменная для хранения данных регистрации

	// Пытаемся распарсить JSON из body в структуру
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Некорректные данные"}`, http.StatusBadRequest)
		return
	}

	// Проверка совпадения паролей
	if req.Password != req.RepeatPassword {
		http.Error(w, `{"error":"Пароли не совпадают"}`, http.StatusBadRequest)
		return
	}

	// Проверка длины пароля
	if len(req.Password) < 8 {
		http.Error(w, `{"error":"Пароль слишком короткий"}`, http.StatusBadRequest)
		return
	}

	// Генерация хэша пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	// Проверка работы bcrypt
	if err != nil {
		http.Error(w, `{"error":"Ошибка сервера"}`, http.StatusInternalServerError)
		return
	}

	// Запрос на вставку пользователя в БД
	_, err = db.DB.Exec(
		"INSERT INTO users (first_name, last_name, email, password_hash) VALUES (?, ?, ?, ?)",
		req.FirstName, req.LastName, req.Email, string(hash),
	)

	// Проверка успешной вставки в БД
	if err != nil {
		log.Println("Ошибка вставки в БД:", err)
		http.Error(w, `{"error":"Email уже используется или другая ошибка"}`, http.StatusBadRequest)
		return
	}

	// Отправка клиенту статуса "ОК"
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
