package main

import (
	"net/http"
	"html/template"

	// "encoding/json"
	// "time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	
	files := []string {
		"templates/pages/main.page.tmpl",
		"templates/layouts/index.layout.tmpl",
	}
	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, nil)
}