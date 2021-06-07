package main

import (
	"net/http"
	"html/template"
	"fmt"
	
	
	// "bytes"
    // "encoding/json"
	// "encoding/json"
	// "time"
)

var mySigningKey = []byte("secret") 

func indexHandler(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Slider []slider
		Hockey []sorevnovanie_and_match
		Volleyball []sorevnovanie_and_match
		Table_tennis []sorevnovanie_and_match
		Field_hockey []sorevnovanie_and_match
		Basketball []sorevnovanie_and_match
		Football []sorevnovanie_and_match
	}

	send := data {
		Slider: Slider(),
		Hockey: Sorev(1),
		Volleyball: Sorev(2),
		Table_tennis: Sorev(3),
		Field_hockey: Sorev(4),
		Basketball: Sorev(5),
		Football: Sorev(6),
	}

	files := []string {
		"templates/pages/main.page.tmpl",
		"templates/layouts/index.layout.tmpl",
		"templates/partials/new.partial.tmpl",
	}
	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, send)
}

func calendarHandler(w http.ResponseWriter, r *http.Request) {
	send := calendar(r.URL.Query().Get("id"))

	files := []string {
		"templates/pages/calendar.page.tmpl",
		"templates/layouts/index.layout.tmpl",
	}
	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, send)
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	send := history(r.URL.Query().Get("id"))
	
	files := []string {
		"templates/pages/history.page.tmpl",
		"templates/layouts/index.layout.tmpl",
	}
	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, send)
}

func matchHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	send := Match(id)

	files := []string {
		"templates/pages/match.page.tmpl",
		"templates/layouts/index.layout.tmpl",
	}
	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, send)
}

func commandsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	send := Commands(id)

	files := []string {
		"templates/pages/commands.page.tmpl",
		"templates/layouts/index.layout.tmpl",
	}
	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, send)
}

func sorevnovanieHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	send := Sorevnivania(id)

	files := []string {
		"templates/pages/sorevnovanie.page.tmpl",
		"templates/layouts/index.layout.tmpl",
	}
	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, send)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var send error
	
	if r.Method == http.MethodPost {
		err := register(w, r)
		if err != nil {
			send = err
		}
	}

	tmpl, _ := template.ParseFiles("templates/register.html")
	tmpl.Execute(w, send)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var send error
	if r.Method == http.MethodPost {
		r.ParseForm()
		id, err := login(r.Form.Get("username"), r.Form.Get("password"))
		if err != nil {
			send = err
		} else {
			token := getToken(mySigningKey, id)
			cookie := http.Cookie{Name: "token", Path: "/", Value: token}
			http.SetCookie(w, &cookie)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

	tmpl, _ := template.ParseFiles("templates/login.html")
	tmpl.Execute(w, send)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	token := w.Header().Get("Authorization")
	claims, err := parseToken(token, mySigningKey)
	if !err {
		fmt.Println("Ошибка токена")
	}
	id := fmt.Sprintf("%v", claims["User_id"])

	send := UserPage(id)
	
	files := []string {
		"templates/pages/user.page.tmpl",
		"templates/layouts/index.layout.tmpl",
	}
	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, send)
}

func logOut(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie {
		Name: "token",
		Path: "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello from middleware")
		cookie, err := r.Cookie("token")
		if err != nil {
			fmt.Println(err)
		} else {
			w.Header().Set("Authorization", cookie.Value)
		}
		next(w, r)
	}
}