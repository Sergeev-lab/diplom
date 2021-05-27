package main

import (
	"net/http"
	"html/template"
	// "encoding/json"
	// "time"
)

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

type data struct {
	Slider []slider
	Hockey []sorev
	Volleyball []sorev
	Table_tennis []sorev
	Field_hockey []sorev
	Basketball []sorev
	Football []sorev
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
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