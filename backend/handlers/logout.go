package handlers

import (
	"net/http"
)

// Обработка выхода из аккаунта
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Разрешение только POST-запросов
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Удаление cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	// Удаление сессии
	cookie, err := r.Cookie("session_id")

	if err == nil {
		delete(sessions, cookie.Value)
	}

	// Ответ клиенту "ОК"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "ok", "message": "Успешный выход"}`))
}
