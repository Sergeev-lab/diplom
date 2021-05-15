package main

import (
	"fmt"
	"net/http"
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Главная страничка
	http.HandleFunc("/", indexHandler)

	// Запуск сервера
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}