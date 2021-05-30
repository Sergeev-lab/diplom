package main

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func main() {

	db, err := sql.Open("mysql", "root:@/diplom")
	if err != nil {
		fmt.Println(err)
	}

	database = db

	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	
	// Главная страничка
	http.HandleFunc("/", indexHandler)

	// Страничка с матчами
	http.HandleFunc("/match/", matchHandler)

	// Страничка с командами
	http.HandleFunc("/commands/", commandsHandler)

	// Страничка с соревнованиями
	http.HandleFunc("/sorevnovanie/", sorevnovanieHandler)

	// Запуск сервера
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}