package main

import (
	"net/http"
	"html/template"
	"fmt"
	
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
	fmt.Println("")
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
	type table struct {
		Sorev string
		Table []tablepoints
	}

	id := r.URL.Query().Get("id")
	send := table {
		Table: Sorevnivania(id),
		Sorev: getSorevName(id),
	}

	files := []string {
		"templates/pages/sorevnovanie.page.tmpl",
		"templates/layouts/index.layout.tmpl",
	}
	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, send)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		user := r.FormValue("username")
		password := r.FormValue("password")
		err := register(user, password)
		if err != nil {
			fmt.Println(err)
		}
	}
	
	tmpl, _ := template.ParseFiles("templates/register.html")
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// token := getToken(mySigningKey)
	// parseToken(token, mySigningKey)

	err := login(r.FormValue("username"), r.FormValue("password"))
	if err != nil {
		fmt.Println(err)
	}

	tmpl, _ := template.ParseFiles("templates/login.html")
	tmpl.Execute(w, nil)
}