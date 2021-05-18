package main

import (
	"fmt"
	"net/http"
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	
	// Главная страничка
	http.HandleFunc("/", indexHandler)

	// Запуск сервера
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}