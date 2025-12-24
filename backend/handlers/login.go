package handlers

import (
	"INside-site/backend/db"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// Структура для принятия данных из JSON-запроса
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Функция для генерации ID сессии
func generateSessionID() (string, error) {
	bytes := make([]byte, 32) // 256 бит
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// Обработка входа в аккаунт
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// preflight-запрос
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Обработка только POST-запросов
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Метод не поддерживается"}`, http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest // Переменная для данных реквеста

	// Парсим данные с полученного JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Некорректный JSON"}`, http.StatusBadRequest)
		return
	}

	var storedHash string // Переменная для хранения хэша пароля

	// Получение хэша пароля пользователя по его email
	err := db.DB.QueryRow("SELECT password_hash FROM users WHERE email = ?", req.Email).Scan(&storedHash)

	// Если email нет в базе
	if err == sql.ErrNoRows {
		http.Error(w, `{"error":"Неверный email или пароль"}`, http.StatusUnauthorized)
		return
	}

	// Если произошла другая ошибка с базой
	if err != nil {
		http.Error(w, `{"error":"Ошибка сервера"}`, http.StatusInternalServerError)
		return
	}

	// Проверка совпадения хэшей паролей
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password))

	// Если не совпадают
	if err != nil {
		http.Error(w, `{"error":"Неверный email или пароль"}`, http.StatusUnauthorized)
		return
	}

	// Генерация сессии
	sessionID, err := generateSessionID()

	// Проверка на успешную генерацию
	if err != nil {
		http.Error(w, `{"error":"Ошибка сервера"}`, http.StatusInternalServerError)
		return
	}

	// Устанавливаем email как ID сессии
	sessions[sessionID] = req.Email

	// Устанавливаем cookie для сесии
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   86400,
	})

	// Отправка клиенту, что всё "ОК"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Успешный вход",
	})
}
