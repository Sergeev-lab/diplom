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
	http.HandleFunc("/", middleware(indexHandler))

	// Страничка с матчами
	http.HandleFunc("/match/", middleware(matchHandler))

	// Страничка с календарем
	http.HandleFunc("/calendar/", calendarHandler)
	// Страничка с историей
	// http.HandleFunc("/history/", historyHandler)

	// Страничка с командами
	http.HandleFunc("/commands/", middleware(commandsHandler))

	// Страничка с соревнованиями
	http.HandleFunc("/sorevnovanie/", middleware(sorevnovanieHandler))

	// Страничка с регистрацией
	http.HandleFunc("/register/", registerHandler)
	// Страничка с входом
	http.HandleFunc("/login/", loginHandler)
	// Страничка личного кабинета
	http.HandleFunc("/user/", middleware(userHandler))
	// Страничка личного кабинета
	http.HandleFunc("/user/logout/", logOut)

	// Запуск сервера
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}