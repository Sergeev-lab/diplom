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

type player struct {
	Number string
	Name string
	Position string
}

type formatch struct {
	Fc_id string
	Fc_name string
	Fc_present string
	Fc_logo string
	Fc_players []player
	Sc_id string
	Sc_name string
	Sc_present string
	Sc_logo string
	Sc_players []player
	Sorev_id string
	Sorev_name string
	Total string
	City string
	Stad string
	Data string
}

type sorev struct {
	Id string
	Name string
	Match []math
}

type math struct {
	Id string
	Fc string
	Sc string
	Total string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Slider []slider
		Hockey []sorev
		Volleyball []sorev
		Table_tennis []sorev
		Field_hockey []sorev
		Basketball []sorev
		Football []sorev
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
	
	type data struct {
		Sorev sorevnovanie
		Table []tablepoints
		Players []commands
	}

	send := data {
		Table: Sorevnivania(id),
		Sorev: getSorev(id),
		Players: getPlayers(id),
	}

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
	type sent struct {
		Data user
		Dost rezults_command
	}

	token := w.Header().Get("Authorization")
	claims, err := parseToken(token, mySigningKey)
	if !err {
		fmt.Println("Ошибка токена")
	}
	id := fmt.Sprintf("%v", claims["User_id"])
	
	dataUser := getUser(id)

	send := sent {
		Data: dataUser,
		Dost: getDost(dataUser.Command.Id),
	}

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(32 << 20)
		fmt.Println(r.Form)
		
	}
	
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