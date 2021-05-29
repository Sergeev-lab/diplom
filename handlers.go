package main

import (
	"net/http"
	"html/template"
	// "fmt"
	// "encoding/json"
	// "time"
)

type formatch struct {
	Fc_id string
	Fc_name string
	Fc_present string
	Fc_logo string
	Sc_id string
	Sc_name string
	Sc_present string
	Sc_logo string
	Sorev_id string
	Sorev_name string
	Total string
	City string
	Stad string
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

func matchHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	send := Match(id)

	tmpl, _ := template.ParseFiles("templates/match.html")
	tmpl.Execute(w, send)
}

func commandsHandler(w http.ResponseWriter, r *http.Request) {
	type info struct {
		City string
		Name string
		Logo string
		Sport string
	}

	type result struct {
		Match_id string
		Fc_id string
		Fc_name string
		Fc_present string
		Fc_logo string
		Sc_id string
		Sc_name string
		Sc_present string
		Sc_logo string
		Stad string
		Total string
		Data string
	}

	type kalendar struct {

	}

	type data struct {
		Info info
		Results result 
		Kalendar kalendar
	}

	// id := r.URL.Query().Get("id")

	send := data {
		Info: info {
			City: "казань" ,
			Name: "динамо-акбарс",
			Logo: "/img/commands/dynamo-kzn.png",
			Sport: "Хоккей на траве",
		},
		Results: result {
			Match_id: "4",
			Fc_id: "3",
			Fc_name: "динамо-акбарс",
			Fc_logo: "/img/commands/dynamo-kzn.png",
			Fc_present: "Казань",
			Sc_id: "2",
			Sc_name: "динамо-электросталь",
			Sc_logo: "/img/commands/elstal.jpg",
			Sc_present: "мос область",
			Stad: "центр хоккея на траве",
			Total: "2:1",
			Data: "14 апреля 2021",
		},
		Kalendar: kalendar {

		},
	}

	tmpl, _ := template.ParseFiles("templates/commands.html")
	tmpl.Execute(w, send)
}