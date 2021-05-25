package main

import (
	"net/http"
	"html/template"
	// "fmt"
	// "encoding/json"
	// "time"
)

type data struct {
	Slider []slider
	Field_hockey []matches
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	send := data {
		Slider: Slider(),
		Field_hockey: Match(),
	}

	files := []string {
		"templates/pages/main.page.tmpl",
		"templates/layouts/index.layout.tmpl",
		"templates/partials/new.partial.tmpl",
	}
	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, send)
}