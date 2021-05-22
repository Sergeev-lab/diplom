package main

import (
	"net/http"
	"html/template"
	// "fmt"
	// "encoding/json"
	// "time"
)

type slider struct {
	Id string
	Img string
	Title string
	Subtitle string
	Description string
	Btn string
}

type match struct {
	Id string
	Name string
	Total string
}

type sorevnovanie struct {
	Id string
	Name string
	Icon string
	Match match 
}

type hockey struct {
	
}

type volleyball struct {
	
}

type tabletennis struct {
	
}

type feildhockey struct {
	Header string
	Href string
	Sorevnovanie sorevnovanie
}

type basketball struct {
	
}

type football struct {
	
}

type content struct {
	Hockey hockey
	Volleyball volleyball
	TableTennis tabletennis
	FieldHockey feildhockey
	Basketball basketball
	Football football
}

type data struct {
	Slider []slider
	Content content
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	
	send := data {
		Slider: Slider(),
		Content: content {
			FieldHockey: feildhockey {
				Header: "Хоккей на траве",
				Href: "/field_hockey/",
				Sorevnovanie: sorevnovanie {
					Id: "3",
					Name: "Кубок России",
					Icon: "/img/sport-icons/field-hockey.svg",
					Match: match {
						Id: "1",
						Name: "Зеленодольск - Казань",
						Total: "1 - 0",
					},
				},
			},
		},
	}

	files := []string {
		"templates/pages/main.page.tmpl",
		"templates/layouts/index.layout.tmpl",
		"templates/partials/new.partial.tmpl",
	}
	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, send)
}