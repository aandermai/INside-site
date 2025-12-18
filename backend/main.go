package main

import (
	"INside-site/backend/db"
	"INside-site/backend/handlers"
	"log"
	"net/http"
)

func main() {
	db.InitDB()
	log.Println("База данных подключена")

	// API
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/profile", handlers.ProfileHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)

	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)

	log.Println("Сервер запущен: http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
