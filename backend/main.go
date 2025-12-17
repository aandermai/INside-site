package main

import (
	"INside-site/backend/db"
	"INside-site/backend/handlers"
	"log"
	"net/http"
)

func main() {
	// Подключение к БД
	db.InitDB()
	log.Println("База данных подключена")

	// Запросы к хендлерам
	http.HandleFunc("/login", handlers.LoginHandler)       // вход
	http.HandleFunc("/register", handlers.RegisterHandler) // регистрация
	http.HandleFunc("/profile", handlers.ProfileHandler)   // данные профиля
	http.HandleFunc("/logout", handlers.LogoutHandler)     // выход

	// Раздача всех файлов из папки frontend
	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	// Запуск сервера
	log.Println("Сервер запущен! Открой в браузере: http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
