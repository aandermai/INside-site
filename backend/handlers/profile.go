package handlers

import (
	"INside-site/backend/db"
	"encoding/json"
	"net/http"
)

var sessions = map[string]string{} // Сессии пользователей (session_id → email)

// Обработка отображения данных профиля
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// preflight-запрос
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Получаем cookie для сессии
	cookie, err := r.Cookie("session_id")

	// Если пользователь не авторизован
	if err != nil {
		http.Error(w, "Не авторизован", http.StatusUnauthorized)
		return
	}

	// Получаем email по сессии пользователя
	email, ok := sessions[cookie.Value]

	// Если нет сессии
	if !ok {
		http.Error(w, "Не авторизован", http.StatusUnauthorized)
		return
	}

	var firstName, lastName string // Переменные для имени и фамилии

	// Запрос к БД, где мы ищем пользователей с подходящим email и берём его имя и фамилию
	err = db.DB.QueryRow("SELECT first_name, last_name FROM users WHERE email = ?", email).
		Scan(&firstName, &lastName)

	// Проверка на успешное выполнение запроса к БД
	if err != nil {
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	// Отправка клиенту имени, фамилии и почты
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"first_name": firstName,
		"last_name":  lastName,
		"email":      email,
	})
}
