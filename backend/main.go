package main

import (
	"backend/db"
	"backend/handlers"
	"log"
	"net/http"
)

func main() {
	db.InitDB()

	http.HandleFunc("/register", handlers.RegisterHandler)

	log.Println("Сервер запущен на порту 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
